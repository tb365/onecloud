package aws

import (
	"yunion.io/x/jsonutils"
)

type SElbBackend struct {

}

func (self *SElbBackend) GetId() string {
	panic("implement me")
}

func (self *SElbBackend) GetName() string {
	panic("implement me")
}

func (self *SElbBackend) GetGlobalId() string {
	panic("implement me")
}

func (self *SElbBackend) GetStatus() string {
	panic("implement me")
}

func (self *SElbBackend) Refresh() error {
	panic("implement me")
}

func (self *SElbBackend) IsEmulated() bool {
	panic("implement me")
}

func (self *SElbBackend) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElbBackend) GetProjectId() string {
	panic("implement me")
}

func (self *SElbBackend) GetWeight() int {
	panic("implement me")
}

func (self *SElbBackend) GetPort() int {
	panic("implement me")
}

func (self *SElbBackend) GetBackendType() string {
	panic("implement me")
}

func (self *SElbBackend) GetBackendRole() string {
	panic("implement me")
}

func (self *SElbBackend) GetBackendId() string {
	panic("implement me")
}

func (self *SElbBackend) SyncConf(port, weight int) error {
	panic("implement me")
}
