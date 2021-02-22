package tasks

import (
	"context"
	"fmt"
	"sort"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/util/netutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/lockman"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/taskman"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/logclient"
	"yunion.io/x/onecloud/pkg/util/rbacutils"
)

type NetworksUnderWireMergeTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(NetworksUnderWireMergeTask{})
}

func (self *NetworksUnderWireMergeTask) taskFailed(ctx context.Context, wire *models.SWire, desc string, err error) {
	d := jsonutils.NewDict()
	d.Set("description", jsonutils.NewString(desc))
	if err != nil {
		d.Set("error", jsonutils.NewString(err.Error()))
	}
	wire.SetStatus(self.UserCred, api.WIRE_STATUS_MERGE_NETWORK_FAILED, d.PrettyString())
	db.OpsLog.LogEvent(wire, db.ACT_MERGE_NETWORK_FAILED, d, self.UserCred)
	logclient.AddActionLogWithStartable(self, wire, logclient.ACT_MERGE_NETWORK, d, self.UserCred, false)
	self.SetStageFailed(ctx, nil)
}

func (self *NetworksUnderWireMergeTask) taskSuccess(ctx context.Context, wire *models.SWire, desc string) {
	d := jsonutils.NewString(desc)
	wire.SetStatus(self.UserCred, api.WIRE_STATUS_AVAILABLE, "")
	db.OpsLog.LogEvent(wire, db.ACT_MERGE_NETWORK, d, self.UserCred)
	logclient.AddActionLogWithStartable(self, wire, logclient.ACT_MERGE_NETWORK, d, self.UserCred, true)
	self.SetStageComplete(ctx, nil)
}

type Net struct {
	*models.SNetwork
	StartIp netutils.IPV4Addr
}

func (self *NetworksUnderWireMergeTask) OnInit(ctx context.Context, obj db.IStandaloneModel, body jsonutils.JSONObject) {
	w := obj.(*models.SWire)
	w.SetStatus(self.UserCred, api.WIRE_STATUS_MERGE_NETWORK, "")

	lockman.LockClass(ctx, models.NetworkManager, db.GetLockClassKey(models.NetworkManager, self.UserCred))
	defer lockman.ReleaseClass(ctx, models.NetworkManager, db.GetLockClassKey(models.NetworkManager, self.UserCred))
	networks, err := w.GetNetworks(self.UserCred, rbacutils.ScopeDomain)
	if err != nil {
		self.taskFailed(ctx, w, "unable to GetNetworks", err)
		return
	}
	if len(networks) <= 1 {
		self.taskSuccess(ctx, w, fmt.Sprintf("num of networks under wire is %d", len(networks)))
	}
	nets := make([]Net, len(networks))
	for i := range nets {
		startIp, _ := netutils.NewIPV4Addr(networks[i].GuestIpStart)
		nets[i] = Net{
			SNetwork: &networks[i],
			StartIp:  startIp,
		}
	}
	sort.Slice(nets, func(i, j int) bool {
		if nets[i].VlanId == nets[j].VlanId {
			return nets[i].StartIp < nets[j].StartIp
		}
		return nets[i].VlanId < nets[j].VlanId
	})
	log.Infof("nets sorted: %s", jsonutils.Marshal(nets))
	for i := 0; i < len(nets)-1; i++ {
		if nets[i].VlanId != nets[i+1].VlanId {
			continue
		}
		// preparenets
		wireNets := make([]*models.SNetwork, 0, len(nets)-2)
		for j := range nets {
			if j != i && j != i+1 {
				wireNets = append(wireNets, nets[i].SNetwork)
			}
		}
		startIp, endIp, err := nets[i].CheckInvalidToMerge(ctx, nets[i+1].SNetwork, wireNets)
		if err != nil {
			log.Debugf("unable to merge network %q to %q: %v", nets[i].GetId(), nets[i+1].GetId(), err)
			continue
		}
		err = nets[i].MergeToNetworkAfterCheck(ctx, self.UserCred, nets[i+1].SNetwork, startIp, endIp)
		if err != nil {
			self.taskFailed(ctx, w, fmt.Sprintf("unable to merge network %q to %q", nets[i].GetId(), nets[i+1].GetId()), err)
			return
		}
	}
	self.taskSuccess(ctx, w, "")
}
