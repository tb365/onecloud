package tasks

import (
	"context"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/taskman"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/logclient"
)

type CloudProviderSyncSkusTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(CloudProviderSyncSkusTask{})
}

func (self *CloudProviderSyncSkusTask) taskFailed(ctx context.Context, provider *models.SCloudprovider, err error) {
	provider.SetStatus(self.UserCred, api.CLOUD_PROVIDER_SYNC_STATUS_ERROR, err.Error())
	db.OpsLog.LogEvent(provider, db.ACT_SYNC_CLOUD_SKUS, err.Error(), self.GetUserCred())
	logclient.AddActionLogWithStartable(self, provider, logclient.ACT_CLOUD_SYNC, err.Error(), self.UserCred, false)
	self.SetStageFailed(ctx, err.Error())
}

func (self *CloudProviderSyncSkusTask) OnInit(ctx context.Context, obj db.IStandaloneModel, body jsonutils.JSONObject) {
	provider := obj.(*models.SCloudprovider)
	regions, err := models.CloudregionManager.GetRegionByProvider(provider.GetId())
	if err != nil {
		self.taskFailed(ctx, provider, err)
	}

	res, _ := body.GetString("resource")
	switch res {
	case models.ServerSkuManager.Keyword():
		for _, region := range regions {
			if err := models.SyncServerSkusByRegion(ctx, self.GetUserCred(), &region); err != nil {
				self.taskFailed(ctx, provider, err)
				return
			}
		}
	case models.ElasticcacheSkuManager.Keyword():
		for _, region := range regions {
			if err := models.SyncElasticCacheSkusByRegion(ctx, self.GetUserCred(), &region); err != nil {
				self.taskFailed(ctx, provider, err)
				return
			}
		}
	case models.ElasticcacheSkuManager.Keyword():
		for _, region := range regions {
			models.SyncRegionDBInstanceSkus(ctx, self.GetUserCred(), region.GetId(), false)
			return
		}
	}

	self.SetStageComplete(ctx, nil)
}
