package aws

import (
	"fmt"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/jsonutils"
)

type SElbBackend struct {
	region *SRegion
	group  *SElbBackendGroup

	Target       Target       `json:"Target"`
	TargetHealth TargetHealth `json:"TargetHealth"`
}

type Target struct {
	ID   string `json:"Id"`
	Port int  `json:"Port"`
}

type TargetHealth struct {
	State       string `json:"State"`
	Reason      string `json:"Reason"`
	Description string `json:"Description"`
}

func (self *SElbBackend) GetId() string {
	return fmt.Sprintf("%s::%s::%s", self.group.GetId(), self.Target.ID, self.Target.Port)
}

func (self *SElbBackend) GetName() string {
	return self.GetId()
}

func (self *SElbBackend) GetGlobalId() string {
	return self.GetId()
}

func (self *SElbBackend) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SElbBackend) Refresh() error {
	// todo: implement me
	return nil
}

func (self *SElbBackend) IsEmulated() bool {
	return false
}

func (self *SElbBackend) GetMetadata() *jsonutils.JSONDict {
	return jsonutils.NewDict()
}

func (self *SElbBackend) GetProjectId() string {
	return ""
}

func (self *SElbBackend) GetWeight() int {
	return 0
}

func (self *SElbBackend) GetPort() int {
	return self.Target.Port
}

func (self *SElbBackend) GetBackendType() string {
	return api.LB_BACKEND_GUEST
}

func (self *SElbBackend) GetBackendRole() string {
	return api.LB_BACKEND_ROLE_DEFAULT
}

func (self *SElbBackend) GetBackendId() string {
	return self.Target.ID
}

func (self *SElbBackend) SyncConf(port, weight int) error {
	panic("implement me")
}
