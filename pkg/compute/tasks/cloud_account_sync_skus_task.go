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

type CloudAccountSyncSkusTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(CloudAccountSyncSkusTask{})
}

func (self *CloudAccountSyncSkusTask) taskFailed(ctx context.Context, account *models.SCloudaccount, err error) {
	account.SetStatus(self.UserCred, api.CLOUD_PROVIDER_SYNC_STATUS_ERROR, err.Error())
	db.OpsLog.LogEvent(account, db.ACT_SYNC_CLOUD_SKUS, err.Error(), self.GetUserCred())
	logclient.AddActionLogWithStartable(self, account, logclient.ACT_CLOUD_SYNC, err.Error(), self.UserCred, false)
	self.SetStageFailed(ctx, err.Error())
}

func (self *CloudAccountSyncSkusTask) OnInit(ctx context.Context, obj db.IStandaloneModel, body jsonutils.JSONObject) {
	account := obj.(*models.SCloudaccount)

	regions := []models.SCloudregion{}
	if regionId, _ := self.GetParams().GetString("cloudregion_id"); len(regionId) > 0 {
		_region, err := db.FetchById(models.CloudregionManager, regionId)
		if err != nil {
			self.taskFailed(ctx, account, err)
			return
		}

		region := _region.(*models.SCloudregion)
		regions = append(regions, *region)
	} else if providerId, _ := self.GetParams().GetString("cloudprovider_id"); len(providerId) > 0 {
		provider, err := db.FetchById(models.CloudproviderManager, providerId)
		if err != nil {
			self.taskFailed(ctx, account, err)
			return
		}

		_regions := provider.(*models.SCloudprovider).GetCloudproviderRegions()
		for i := range _regions {
			region := _regions[i].GetRegion()
			regions = append(regions, *region)
		}
	} else {
		providers := account.GetEnabledCloudproviders()
		for _, provider := range providers {
			_regions := provider.GetCloudproviderRegions()
			for i := range _regions {
				region := _regions[i].GetRegion()
				regions = append(regions, *region)
			}
		}
	}

	res, _ := self.GetParams().GetString("resource")
	for _, region := range regions {
		switch res {
		case models.ServerSkuManager.Keyword():
			if err := models.SyncServerSkusByRegion(ctx, self.GetUserCred(), &region); err != nil {
				self.taskFailed(ctx, account, err)
				return
			}
		case models.ElasticcacheSkuManager.Keyword():
			if err := models.SyncElasticCacheSkusByRegion(ctx, self.GetUserCred(), &region); err != nil {
				self.taskFailed(ctx, account, err)
				return
			}
		case models.ElasticcacheSkuManager.Keyword():
			models.SyncRegionDBInstanceSkus(ctx, self.GetUserCred(), region.GetId(), false)
		}
	}

	self.SetStageComplete(ctx, nil)
}
