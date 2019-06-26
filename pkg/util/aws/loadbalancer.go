package aws

import (
	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

/*
https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/Welcome.html
*/

type SElb struct {
	Type                  string             `json:"Type"`
	Scheme                string             `json:"Scheme"`
	IPAddressType         string             `json:"IpAddressType"`
	VpcID                 string             `json:"VpcId"`
	AvailabilityZones     []AvailabilityZone `json:"AvailabilityZones"`
	CreatedTime           string             `json:"CreatedTime"`
	CanonicalHostedZoneID string             `json:"CanonicalHostedZoneId"`
	DNSName               string             `json:"DNSName"`
	SecurityGroups        []string           `json:"SecurityGroups"`
	LoadBalancerName      string             `json:"LoadBalancerName"`
	State                 State              `json:"State"`
	LoadBalancerArn       string             `json:"LoadBalancerArn"`
}

type AvailabilityZone struct {
	LoadBalancerAddresses []LoadBalancerAddress `json:"LoadBalancerAddresses"`
	ZoneName              string                `json:"ZoneName"`
	SubnetID              string                `json:"SubnetId"`
}

type LoadBalancerAddress struct {
	IPAddress    string `json:"IpAddress"`
	AllocationID string `json:"AllocationId"`
}

type State struct {
	Code string `json:"Code"`
}

func (self *SElb) GetId() string {
	return self.LoadBalancerArn
}

func (self *SElb) GetName() string {
	return self.LoadBalancerName
}

func (self *SElb) GetGlobalId() string {
	return self.GetId()
}

func (self *SElb) GetStatus() string {
	switch self.State.Code {
	case "provisioning":
		return api.LB_STATUS_INIT
	case "active":
		return api.LB_STATUS_ENABLED
	case "failed":
		return api.LB_STATUS_START_FAILED
	default:
		return api.LB_STATUS_UNKNOWN
	}
}

func (self *SElb) Refresh() error {
	panic("implement me")
}

func (self *SElb) IsEmulated() bool {
	return false
}

func (self *SElb) GetMetadata() *jsonutils.JSONDict {
	return jsonutils.NewDict()
}

func (self *SElb) GetProjectId() string {
	return ""
}

func (self *SElb) GetAddress() string {
	// todo: dns?
	return self.DNSName
}

func (self *SElb) GetAddressType() string {
	switch self.Scheme {
	case "internal":
		return api.LB_ADDR_TYPE_INTRANET
	case "internet-facing":
		return api.LB_ADDR_TYPE_INTERNET
	default:
		return api.LB_ADDR_TYPE_INTRANET
	}
}

func (self *SElb) GetNetworkType() string {
	return api.LB_NETWORK_TYPE_VPC
}

// todo: 过个network id怎么兼容？
func (self *SElb) GetNetworkId() string {
	return self.AvailabilityZones[0].SubnetID
}

func (self *SElb) GetVpcId() string {
	return self.VpcID
}

// todo: 过个network id怎么兼容？
func (self *SElb) GetZoneId() string {
	return self.AvailabilityZones[0].ZoneName
}

func (self *SElb) GetLoadbalancerSpec() string {
	return ""
}

func (self *SElb) GetChargeType() string {
	return api.LB_CHARGE_TYPE_BY_HOUR
}

func (self *SElb) GetEgressMbps() int {
	return 0
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