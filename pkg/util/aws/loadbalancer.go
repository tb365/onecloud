package aws

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

/*
https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/Welcome.html
*/

type SElb struct {

}

func (self *SElb) GetId() string {
	panic("implement me")
}

func (self *SElb) GetName() string {
	panic("implement me")
}

func (self *SElb) GetGlobalId() string {
	panic("implement me")
}

func (self *SElb) GetStatus() string {
	panic("implement me")
}

func (self *SElb) Refresh() error {
	panic("implement me")
}

func (self *SElb) IsEmulated() bool {
	panic("implement me")
}

func (self *SElb) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElb) GetProjectId() string {
	panic("implement me")
}

func (self *SElb) GetAddress() string {
	panic("implement me")
}

func (self *SElb) GetAddressType() string {
	panic("implement me")
}

func (self *SElb) GetNetworkType() string {
	panic("implement me")
}

func (self *SElb) GetNetworkId() string {
	panic("implement me")
}

func (self *SElb) GetVpcId() string {
	panic("implement me")
}

func (self *SElb) GetZoneId() string {
	panic("implement me")
}

func (self *SElb) GetLoadbalancerSpec() string {
	panic("implement me")
}

func (self *SElb) GetChargeType() string {
	panic("implement me")
}

func (self *SElb) GetEgressMbps() int {
	panic("implement me")
}

func (self *SElb) Delete() error {
	panic("implement me")
}

func (self *SElb) Start() error {
	panic("implement me")
}

func (self *SElb) Stop() error {
	panic("implement me")
}

func (self *SElb) GetILoadBalancerListeners() ([]cloudprovider.ICloudLoadbalancerListener, error) {
	panic("implement me")
}

func (self *SElb) GetILoadBalancerBackendGroups() ([]cloudprovider.ICloudLoadbalancerBackendGroup, error) {
	panic("implement me")
}

func (self *SElb) CreateILoadBalancerBackendGroup(group *cloudprovider.SLoadbalancerBackendGroup) (cloudprovider.ICloudLoadbalancerBackendGroup, error) {
	panic("implement me")
}

func (self *SElb) GetILoadBalancerBackendGroupById(groupId string) (cloudprovider.ICloudLoadbalancerBackendGroup, error) {
	panic("implement me")
}

func (self *SElb) CreateILoadBalancerListener(listener *cloudprovider.SLoadbalancerListener) (cloudprovider.ICloudLoadbalancerListener, error) {
	panic("implement me")
}

func (self *SElb) GetILoadBalancerListenerById(listenerId string) (cloudprovider.ICloudLoadbalancerListener, error) {
	panic("implement me")
}
