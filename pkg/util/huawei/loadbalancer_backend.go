package huawei

import (
	"fmt"
	"yunion.io/x/jsonutils"
	api "yunion.io/x/onecloud/pkg/apis/compute"
)

type SElbBackend struct {
	region       *SRegion
	lb           *SLoadbalancer
	backendGroup *SElbBackendGroup

	Name            string `json:"name"`
	Weight          int    `json:"weight"`
	AdminStateUp    bool   `json:"admin_state_up"`
	SubnetID        string `json:"subnet_id"`
	TenantID        string `json:"tenant_id"`
	ProjectID       string `json:"project_id"`
	Address         string `json:"address"`
	ProtocolPort    int    `json:"protocol_port"`
	OperatingStatus string `json:"operating_status"`
	ID              string `json:"id"`
}

func (self *SElbBackend) GetId() string {
	return self.ID
}

func (self *SElbBackend) GetName() string {
	return self.Name
}

func (self *SElbBackend) GetGlobalId() string {
	return self.GetId()
}

func (self *SElbBackend) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SElbBackend) Refresh() error {
	m := self.lb.region.ecsClient.ElbBackend
	err := m.SetBackendGroupId(self.backendGroup.GetId())
	if err != nil {
		return err
	}

	backend := SElbBackend{}
	err = DoGet(m.Get, self.GetId(), nil, &backend)
	if err != nil {
		return err
	}

	backend.lb = self.lb
	backend.backendGroup = self.backendGroup
	err = jsonutils.Update(self, backend)
	if err != nil {
		return err
	}

	return nil
}

func (self *SElbBackend) IsEmulated() bool {
	return false
}

func (self *SElbBackend) GetMetadata() *jsonutils.JSONDict {
	return nil
}

func (self *SElbBackend) GetProjectId() string {
	return ""
}

func (self *SElbBackend) GetWeight() int {
	return self.Weight
}

func (self *SElbBackend) GetPort() int {
	return self.ProtocolPort
}

func (self *SElbBackend) GetBackendType() string {
	return api.LB_BACKEND_GUEST
}

func (self *SElbBackend) GetBackendRole() string {
	return api.LB_BACKEND_ROLE_DEFAULT
}

func (self *SElbBackend) GetBackendId() string {
	// todo: 华为backendId fix
	return ""
}

func (self *SElbBackend) SyncConf(port, weight int) error {
	if port > 0 {
		return fmt.Errorf("Elb backend SyncConf unsupport modify port")
	}

	params := jsonutils.NewDict()
	memberObj := jsonutils.NewDict()
	memberObj.Set("weight", jsonutils.NewInt(int64(weight)))
	params.Set("member", memberObj)
	err := self.lb.region.ecsClient.ElbBackend.SetBackendGroupId(self.backendGroup.GetId())
	if err != nil {
		return err
	}
	return DoUpdate(self.lb.region.ecsClient.ElbBackend.Update, self.GetId(), params, nil)
}
