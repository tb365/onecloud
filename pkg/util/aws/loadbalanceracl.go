package aws

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SElbACL struct {
}

func (self *SElbACL) GetId() string {
	panic("implement me")
}

func (self *SElbACL) GetName() string {
	panic("implement me")
}

func (self *SElbACL) GetGlobalId() string {
	panic("implement me")
}

func (self *SElbACL) GetStatus() string {
	panic("implement me")
}

func (self *SElbACL) Refresh() error {
	panic("implement me")
}

func (self *SElbACL) IsEmulated() bool {
	panic("implement me")
}

func (self *SElbACL) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElbACL) GetProjectId() string {
	panic("implement me")
}

func (self *SElbACL) GetAclListenerID() string {
	panic("implement me")
}

func (self *SElbACL) GetAclEntries() []cloudprovider.SLoadbalancerAccessControlListEntry {
	panic("implement me")
}

func (self *SElbACL) Sync(acl *cloudprovider.SLoadbalancerAccessControlList) error {
	panic("implement me")
}

func (self *SElbACL) Delete() error {
	panic("implement me")
}
