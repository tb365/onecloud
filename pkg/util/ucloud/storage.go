package ucloud

import (
	"fmt"
	"strings"

	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/cloudprovider"
	"yunion.io/x/onecloud/pkg/compute/models"
)

type SStorage struct {
	zone        *SZone
	storageType string
}

func (self *SStorage) GetId() string {
	return fmt.Sprintf("%s-%s-%s", self.zone.region.client.providerId, self.zone.GetId(), self.storageType)
}

func (self *SStorage) GetName() string {
	return fmt.Sprintf("%s-%s-%s", self.zone.region.client.providerName, self.zone.GetId(), self.storageType)
}

func (self *SStorage) GetGlobalId() string {
	return fmt.Sprintf("%s-%s-%s", self.zone.region.client.providerId, self.zone.GetGlobalId(), self.storageType)
}

func (self *SStorage) GetStatus() string {
	return models.STORAGE_ONLINE
}

func (self *SStorage) Refresh() error {
	return nil
}

func (self *SStorage) IsEmulated() bool {
	return true
}

func (self *SStorage) GetMetadata() *jsonutils.JSONDict {
	return nil
}

func (self *SStorage) GetIStoragecache() cloudprovider.ICloudStoragecache {
	return self.zone.region.getStoragecache()
}

func (self *SStorage) GetIZone() cloudprovider.ICloudZone {
	return self.zone
}

func (self *SStorage) GetIDisks() ([]cloudprovider.ICloudDisk, error) {
	disks, err := self.zone.region.GetDisks(self.zone.GetId(), "", nil)
	if err != nil {
		return nil, err
	}

	filtedDisks := make([]SDisk, 0)
	for _, disk := range disks {
		// ssd 盘
		if self.storageType == models.STORAGE_UCLOUD_CLOUD_SSD && strings.Contains(disk.DiskType, "SSD") {
			filtedDisks = append(filtedDisks, disk)
		}

		// 普通盘
		if self.storageType == models.STORAGE_UCLOUD_CLOUD_NORMAL && !strings.Contains(disk.DiskType, "SSD") {
			filtedDisks = append(filtedDisks, disk)
		}
	}

	idisks := make([]cloudprovider.ICloudDisk, len(filtedDisks))
	for i := 0; i < len(filtedDisks); i += 1 {
		filtedDisks[i].storage = self
		idisks[i] = &filtedDisks[i]
	}
	return idisks, nil
}

func (self *SStorage) GetStorageType() string {
	return self.storageType
}

func (self *SStorage) GetMediumType() string {
	if self.storageType == models.STORAGE_UCLOUD_CLOUD_SSD {
		return models.DISK_TYPE_SSD
	} else {
		return models.DISK_TYPE_ROTATE
	}
}

func (self *SStorage) GetCapacityMB() int {
	return 0 // unlimited
}

func (self *SStorage) GetStorageConf() jsonutils.JSONObject {
	return jsonutils.NewDict()
}

func (self *SStorage) GetEnabled() bool {
	return true
}

func (self *SStorage) GetManagerId() string {
	return self.zone.region.client.providerId
}

func (self *SStorage) CreateIDisk(name string, sizeGb int, desc string) (cloudprovider.ICloudDisk, error) {
	diskId, err := self.zone.region.CreateDisk(self.zone.GetId(), self.storageType, name, sizeGb)
	if err != nil {
		return nil, err
	}

	disk, err := self.zone.region.GetDisk(diskId)
	if err != nil {
		return nil, err
	}

	disk.storage = self
	return disk, nil
}

func (self *SStorage) GetIDiskById(idStr string) (cloudprovider.ICloudDisk, error) {
	if disk, err := self.zone.region.GetDisk(idStr); err != nil {
		return nil, err
	} else {
		disk.storage = self
		return disk, nil
	}
}

func (self *SStorage) GetMountPoint() string {
	return ""
}

func (self *SStorage) IsSysDiskStore() bool {
	return true
}
