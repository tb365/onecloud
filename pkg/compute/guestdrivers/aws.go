// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package guestdrivers

import (
	"context"
	"fmt"
	"strings"

	"yunion.io/x/pkg/utils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/taskman"
	"yunion.io/x/onecloud/pkg/cloudprovider"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/util/billing"
)

type SAwsGuestDriver struct {
	SManagedVirtualizedGuestDriver
}

func init() {
	driver := SAwsGuestDriver{}
	models.RegisterGuestDriver(&driver)
}

func fetchAwsUserName(desc cloudprovider.SManagedVMCreateConfig) string {
	// 非公有云官方镜像
	if desc.ImageType != "system" {
		return "root"
	}

	// 公有云官方镜像
	dist := strings.ToLower(desc.OsDistribution)
	if strings.Contains(dist, "centos") {
		return "centos"
	} else if strings.Contains(dist, "ubuntu") {
		return "ubuntu"
	} else if strings.Contains(dist, "windows") {
		return "Administrator"
	} else if strings.Contains(dist, "debian") {
		return "admin"
	} else if strings.Contains(dist, "suse") {
		return "ec2-user"
	} else if strings.Contains(dist, "fedora") {
		return "ec2-user"
	} else if strings.Contains(dist, "rhel") || strings.Contains(dist, "redhat") {
		return "ec2-user"
	} else if strings.Contains(dist, "amazon linux") {
		return "ec2-user"
	} else {
		return "ec2-user"
	}
}

func (self *SAwsGuestDriver) GetLinuxDefaultAccount(desc cloudprovider.SManagedVMCreateConfig) string {
	// return fetchAwsUserName(desc)
	return models.VM_AWS_DEFAULT_LOGIN_USER
}

func (self *SAwsGuestDriver) GetHypervisor() string {
	return models.HYPERVISOR_AWS
}

func (self *SAwsGuestDriver) GetDefaultSysDiskBackend() string {
	return models.STORAGE_GP2_SSD
}

func (self *SAwsGuestDriver) GetMinimalSysDiskSizeGb() int {
	return 10
}

func (self *SAwsGuestDriver) GetStorageTypes() []string {
	return []string{
		models.STORAGE_GP2_SSD,
		models.STORAGE_IO1_SSD,
		models.STORAGE_ST1_HDD,
		models.STORAGE_SC1_HDD,
		models.STORAGE_STANDARD_HDD,
	}
}

func (self *SAwsGuestDriver) ChooseHostStorage(host *models.SHost, backend string) *models.SStorage {
	storages := host.GetAttachedStorages("")
	for i := 0; i < len(storages); i += 1 {
		if storages[i].StorageType == backend {
			return &storages[i]
		}
	}

	for _, stype := range self.GetStorageTypes() {
		for i := 0; i < len(storages); i += 1 {
			if storages[i].StorageType == stype {
				return &storages[i]
			}
		}
	}
	return nil
}

func (self *SAwsGuestDriver) GetDetachDiskStatus() ([]string, error) {
	return []string{models.VM_READY, models.VM_RUNNING}, nil
}

func (self *SAwsGuestDriver) GetAttachDiskStatus() ([]string, error) {
	return []string{models.VM_READY, models.VM_RUNNING}, nil
}

func (self *SAwsGuestDriver) GetRebuildRootStatus() ([]string, error) {
	return []string{models.VM_READY, models.VM_RUNNING}, nil
}

func (self *SAwsGuestDriver) GetChangeConfigStatus() ([]string, error) {
	return []string{models.VM_READY}, nil
}

func (self *SAwsGuestDriver) GetDeployStatus() ([]string, error) {
	return []string{models.VM_READY, models.VM_RUNNING}, nil
}

func (self *SAwsGuestDriver) RequestDetachDisk(ctx context.Context, guest *models.SGuest, task taskman.ITask) error {
	return guest.StartSyncTask(ctx, task.GetUserCred(), false, task.GetTaskId())
}

func (self *SAwsGuestDriver) ValidateCreateData(ctx context.Context, userCred mcclient.TokenCredential, input *api.ServerCreateInput) (*api.ServerCreateInput, error) {
	return self.SManagedVirtualizedGuestDriver.ValidateCreateData(ctx, userCred, input)
}

func (self *SAwsGuestDriver) ValidateResizeDisk(guest *models.SGuest, disk *models.SDisk, storage *models.SStorage) error {
	// https://docs.amazonaws.cn/AWSEC2/latest/UserGuide/stop-start.html
	if !utils.IsInStringArray(guest.Status, []string{models.VM_RUNNING, models.VM_READY}) {
		return fmt.Errorf("Cannot resize disk when guest in status %s", guest.Status)
	}
	if disk.DiskType == models.DISK_TYPE_SYS && !utils.IsInStringArray(storage.StorageType, []string{models.STORAGE_IO1_SSD, models.STORAGE_STANDARD_HDD, models.STORAGE_GP2_SSD}) {
		return fmt.Errorf("Cannot resize system disk with unsupported volumes type %s", storage.StorageType)
	}
	if !utils.IsInStringArray(storage.StorageType, []string{models.STORAGE_GP2_SSD, models.STORAGE_IO1_SSD, models.STORAGE_ST1_HDD, models.STORAGE_SC1_HDD, models.STORAGE_STANDARD_HDD}) {
		return fmt.Errorf("Cannot resize %s disk", storage.StorageType)
	}
	return nil
}

func (self *SAwsGuestDriver) GetGuestInitialStateAfterCreate() string {
	return models.VM_RUNNING
}

func (self *SAwsGuestDriver) GetGuestInitialStateAfterRebuild() string {
	return models.VM_READY
}

/*func (self *SAwsGuestDriver) RequestDeployGuestOnHost(ctx context.Context, guest *models.SGuest, host *models.SHost, task taskman.ITask) error {
	config, err := guest.GetDeployConfigOnHost(ctx, task.GetUserCred(), host, task.GetParams())
	if err != nil {
		log.Errorf("GetDeployConfigOnHost error: %v", err)
		return err
	}
	log.Debugf("RequestDeployGuestOnHost: %s", config)
	desc := cloudprovider.SManagedVMCreateConfig{}
	if err := desc.GetConfig(config); err != nil {
		return err
	}
	action, err := config.GetString("action")
	if err != nil {
		return err
	}

	ihost, err := host.GetIHost()
	if err != nil {
		return err
	}

	username := fetchAwsUserName(desc)

	switch action {
	case "create":
		taskman.LocalTaskRun(task, func() (jsonutils.JSONObject, error) {

			iVM, createErr := ihost.CreateVM(&desc)
			if createErr != nil {
				return nil, createErr
			}

			guest.SetExternalId(task.GetUserCred(), iVM.GetGlobalId())

			log.Debugf("VMcreated %s, wait status running ...", iVM.GetGlobalId())
			err = cloudprovider.WaitStatus(iVM, models.VM_RUNNING, time.Second*5, time.Second*1800)
			if err != nil {
				return nil, err
			}
			log.Debugf("VMcreated %s, and status is ready", iVM.GetGlobalId())

			iVM, err = ihost.GetIVMById(iVM.GetGlobalId())
			if err != nil {
				log.Errorf("cannot find vm %s", err)
				return nil, err
			}

			data := fetchIVMinfo(desc, iVM, guest.Id, username, desc.Password, action)
			return data, nil
		})
	case "deploy":
		iVM, err := ihost.GetIVMById(guest.GetExternalId())
		if err != nil || iVM == nil {
			log.Errorf("cannot find vm %s", err)
			return fmt.Errorf("cannot find vm")
		}

		params := task.GetParams()
		log.Debugf("Deploy VM params %s", params.String())
		deleteKeypair := jsonutils.QueryBoolean(params, "__delete_keypair__", false)

		taskman.LocalTaskRun(task, func() (jsonutils.JSONObject, error) {
			err := iVM.DeployVM(ctx, desc.Name, desc.Password, desc.PublicKey, deleteKeypair, desc.Description)
			if err != nil {
				return nil, err
			}

			data := fetchIVMinfo(desc, iVM, guest.Id, username, desc.Password, action)
			return data, nil
		})
	case "rebuild":
		iVM, err := ihost.GetIVMById(guest.GetExternalId())
		if err != nil || iVM == nil {
			log.Errorf("cannot find vm %s", err)
			return fmt.Errorf("cannot find vm")
		}

		taskman.LocalTaskRun(task, func() (jsonutils.JSONObject, error) {
			diskId, err := iVM.RebuildRoot(ctx, desc.ExternalImageId, desc.Password, desc.PublicKey, desc.SysDisk.SizeGB)
			if err != nil {
				return nil, err
			}

			log.Debugf("VMrebuildRoot %s new diskID %s, wait status ready ...", iVM.GetGlobalId(), diskId)

			err = cloudprovider.WaitStatus(iVM, models.VM_READY, time.Second*5, time.Second*1800)
			if err != nil {
				return nil, err
			}
			log.Debugf("VMrebuildRoot %s, and status is ready", iVM.GetGlobalId())

			maxWaitSecs := 300
			waited := 0

			for {
				// hack, wait disk number consistent
				idisks, err := iVM.GetIDisks()
				if err != nil {
					log.Errorf("fail to find VM idisks %s", err)
					return nil, err
				}
				if len(idisks) < len(desc.DataDisks)+1 {
					if waited > maxWaitSecs {
						log.Errorf("inconsistent disk number, wait timeout, must be something wrong on remote")
						return nil, cloudprovider.ErrTimeout
					}
					log.Debugf("inconsistent disk number???? %d != %d", len(idisks), len(desc.DataDisks)+1)
					time.Sleep(time.Second * 5)
					waited += 5
				} else {
					if idisks[0].GetGlobalId() != diskId {
						log.Errorf("system disk id inconsistent %s != %s", idisks[0].GetGlobalId(), diskId)
						return nil, fmt.Errorf("inconsistent sys disk id after rebuild root")
					}

					break
				}
			}

			data := fetchIVMinfo(desc, iVM, guest.Id, username, desc.Password, action)

			return data, nil
		})
	default:
		log.Errorf("RequestDeployGuestOnHost: Action %s not supported", action)
		return fmt.Errorf("Action %s not supported", action)
	}

	return nil
}*/

func (self *SAwsGuestDriver) IsSupportedBillingCycle(bc billing.SBillingCycle) bool {
	return false
}
