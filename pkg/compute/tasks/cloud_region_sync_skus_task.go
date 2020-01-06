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

type CloudRegionSyncSkusTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(CloudRegionSyncSkusTask{})
}

func (self *CloudRegionSyncSkusTask) taskFailed(ctx context.Context, region *models.SCloudregion, err error) {
	region.SetStatus(self.UserCred, api.CLOUD_PROVIDER_SYNC_STATUS_ERROR, err.Error())
	db.OpsLog.LogEvent(region, db.ACT_SYNC_CLOUD_SKUS, err.Error(), self.GetUserCred())
	logclient.AddActionLogWithStartable(self, region, logclient.ACT_CLOUD_SYNC, err.Error(), self.UserCred, false)
	self.SetStageFailed(ctx, err.Error())
}

func (self *CloudRegionSyncSkusTask) OnInit(ctx context.Context, obj db.IStandaloneModel, body jsonutils.JSONObject) {
	region := obj.(*models.SCloudregion)

	res, _ := body.GetString("resource")
	switch res {
	case models.ServerSkuManager.Keyword():
		if err := models.SyncServerSkusByRegion(ctx, self.GetUserCred(), region); err != nil {
			self.taskFailed(ctx, region, err)
			return
		}
	case models.ElasticcacheSkuManager.Keyword():
		if err := models.SyncElasticCacheSkusByRegion(ctx, self.GetUserCred(), region); err != nil {
			self.taskFailed(ctx, region, err)
			return
		}
	case models.ElasticcacheSkuManager.Keyword():
		models.SyncRegionDBInstanceSkus(ctx, self.GetUserCred(), region.GetId(), false)
		return
	}

	self.SetStageComplete(ctx, nil)
}
