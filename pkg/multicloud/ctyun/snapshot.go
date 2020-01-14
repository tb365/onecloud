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

package ctyun

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"

	api "yunion.io/x/onecloud/pkg/apis/compute"
)

type SSnapshot struct {
	region *SRegion

	Status           string `json:"status"`
	Description      string `json:"description"`
	AvailabilityZone string `json:"availability_zone"`
	VolumeID         string `json:"volume_id"`
	FailReason       string `json:"fail_reason"`
	ID               string `json:"id"`
	Size             int32  `json:"size"`
	Container        string `json:"container"`
	Name             string `json:"name"`
	CreatedAt        string `json:"created_at"`
}

func (self *SSnapshot) GetId() string {
	return self.ID
}

func (self *SSnapshot) GetName() string {
	return self.Name
}

func (self *SSnapshot) GetGlobalId() string {
	return self.GetId()
}

func (self *SSnapshot) GetStatus() string {
	switch self.Status {
	case "available":
		return api.SNAPSHOT_READY
	case "creating":
		return api.SNAPSHOT_CREATING
	case "deleting":
		return api.SNAPSHOT_DELETING
	case "error_deleting", "error":
		return api.SNAPSHOT_FAILED
	case "rollbacking":
		return api.SNAPSHOT_ROLLBACKING
	default:
		return api.SNAPSHOT_UNKNOWN
	}
}

func (self *SSnapshot) Refresh() error {
	snapshot, err := self.region.GetSnapshot(self.VolumeID, self.GetId())
	if err != nil {
		return err
	}

	if err := jsonutils.Update(self, snapshot); err != nil {
		return err
	}

	return nil
}

func (self *SSnapshot) IsEmulated() bool {
	return false
}

func (self *SSnapshot) GetMetadata() *jsonutils.JSONDict {
	return nil
}

func (self *SSnapshot) GetProjectId() string {
	return ""
}

func (self *SSnapshot) GetSizeMb() int32 {
	return self.Size * 1024
}

func (self *SSnapshot) GetDiskId() string {
	return self.VolumeID
}

func (self *SSnapshot) GetDiskType() string {
	disk, err := self.region.GetDisk(self.VolumeID)
	if err != nil {
		log.Debugf("SSnapshot.GetDiskType.GetDisk %s", err)
		return ""
	}

	return disk.GetDiskType()
}

func (self *SSnapshot) Delete() error {
	_, err := self.region.DeleteSnapshot(self.GetId())
	if err != nil {
		return errors.Wrap(err, "Snapshot.Delete.DeleteSnapshot")
	}

	return nil
}

func (self *SRegion) GetSnapshot(diskId string, snapshotId string) (*SSnapshot, error) {
	snapshots, err := self.GetSnapshots(diskId)
	if err != nil {
		return nil, errors.Wrap(err, "SRegion.GetSnapshot.GetSnapshots")
	}

	for i := range snapshots {
		snapshot := snapshots[i]
		if snapshot.ID == snapshotId {
			snapshot.region = self
			return &snapshot, nil
		}
	}

	return nil, errors.Wrap(errors.ErrNotFound, "SRegion.GetSnapshot")
}

func (self *SRegion) GetSnapshots(diskId string) ([]SSnapshot, error) {
	params := map[string]string{
		"regionId": self.GetId(),
	}

	if len(diskId) > 0 {
		params["volumeId"] = diskId
	}

	resp, err := self.client.DoGet("/apiproxy/v3/ondemand/queryVBSDetails", params)
	if err != nil {
		return nil, errors.Wrap(err, "SRegion.GetSnapshots.DoGet")
	}

	ret := make([]SSnapshot, 0)
	err = resp.Unmarshal(&ret, "returnObj", "backups")
	if err != nil {
		return nil, errors.Wrap(err, "SRegion.GetSnapshots.Unmarshal")
	}

	for i := range ret {
		ret[i].region = self
	}

	return ret, nil
}

func (self *SRegion) DeleteSnapshot(vbsId string) (string, error) {
	params := map[string]jsonutils.JSONObject{
		"regionId": jsonutils.NewString(self.GetId()),
		"vbsId":    jsonutils.NewString(vbsId),
	}

	resp, err := self.client.DoPost("/apiproxy/v3/ondemand/deleteVBS", params)
	if err != nil {
		return "", errors.Wrap(err, "SRegion.DeleteSnapshot.DoPost")
	}

	var ok bool
	err = resp.Unmarshal(&ok, "returnObj", "status")
	if !ok {
		msg, _ := resp.GetString("message")
		return "", fmt.Errorf("SRegion.DeleteSnapshot.JobFailed %s", msg)
	}

	var jobId string
	err = resp.Unmarshal(&jobId, "returnObj", "data")
	if err != nil {
		return "", errors.Wrap(err, "SRegion.DeleteSnapshot.Unmarshal")
	}

	return jobId, nil
}
