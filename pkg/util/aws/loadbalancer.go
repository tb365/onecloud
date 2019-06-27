package aws

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

/*
https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/Welcome.html
*/

type SElb struct {
	region *SRegion

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
	return self.region.DeleteElb(self.GetId())
}

func (self *SElb) Start() error {
	return nil
}

func (self *SElb) Stop() error {
	return cloudprovider.ErrNotSupported
}

func (self *SElb) GetILoadBalancerListeners() ([]cloudprovider.ICloudLoadbalancerListener, error) {
	listeners, err := self.region.GetElbListeners(self.GetId())
	if err != nil {
		return nil, err
	}

	ret := make([]cloudprovider.ICloudLoadbalancerListener, len(listeners))
	for i := range listeners {
		listeners[i].lb = self
		ret[i] = &listeners[i]
	}

	return ret, nil
}

func (self *SElb) GetILoadBalancerBackendGroups() ([]cloudprovider.ICloudLoadbalancerBackendGroup, error) {
	backendgroups, err := self.region.GetElbBackendgroups(self.GetId(), nil)
	if err != nil {
		return nil, err
	}

	ibackendgroups := make([]cloudprovider.ICloudLoadbalancerBackendGroup, len(backendgroups))
	for i := range backendgroups {
		backendgroups[i].lb = self
		ibackendgroups[i] = &backendgroups[i]
	}

	return ibackendgroups, nil
}

func (self *SElb) CreateILoadBalancerBackendGroup(group *cloudprovider.SLoadbalancerBackendGroup) (cloudprovider.ICloudLoadbalancerBackendGroup, error) {
	panic("implement me")
}

func (self *SElb) GetILoadBalancerBackendGroupById(groupId string) (cloudprovider.ICloudLoadbalancerBackendGroup, error) {
	return self.region.GetElbBackendgroup(groupId)
}

func (self *SElb) CreateILoadBalancerListener(listener *cloudprovider.SLoadbalancerListener) (cloudprovider.ICloudLoadbalancerListener, error) {
	panic("implement me")
}

func (self *SElb) GetILoadBalancerListenerById(listenerId string) (cloudprovider.ICloudLoadbalancerListener, error) {
	return self.region.GetElbListener(listenerId)
}

func (self *SRegion) DeleteElb(elbId string) error {
	client, err := self.GetElbV2Client()
	if err != nil {
		return err
	}

	params := &elbv2.DeleteLoadBalancerInput{}
	params.SetLoadBalancerArn(elbId)
	_, err = client.DeleteLoadBalancer(params)
	if err != nil {
		return err
	}

	return nil
}

func (self *SRegion) GetElbBackendgroups(elbId string,backendgroupIds []string) ([]SElbBackendGroup, error) {
	client, err := self.GetElbV2Client()
	if err != nil {
		return nil, err
	}

	params := &elbv2.DescribeTargetGroupsInput{}
	params.SetLoadBalancerArn(elbId)
	if len(backendgroupIds) > 0 {
		v := make([]*string, len(backendgroupIds))
		for i := range backendgroupIds {
			v[i] = &backendgroupIds[i]
		}

		params.SetTargetGroupArns(v)
	}

	ret, err := client.DescribeTargetGroups(params)
	if err != nil {
		return nil, err
	}

	backendgroups := []SElbBackendGroup{}
	err = unmarshalAwsOutput(ret.String(), "TargetGroups", backendgroups)
	if err != nil {
		return nil, err
	}

	for i := range backendgroups {
		backendgroups[i].region = self
	}

	return backendgroups, nil
}

func (self *SRegion) GetElbBackendgroup(backendgroupId string) (*SElbBackendGroup, error) {
	client, err := self.GetElbV2Client()
	if err != nil {
		return nil, err
	}

	params := &elbv2.DescribeTargetGroupsInput{}
	params.SetTargetGroupArns([]*string{&backendgroupId})

	ret, err := client.DescribeTargetGroups(params)
	if err != nil {
		return nil, err
	}

	backendgroups := []SElbBackendGroup{}
	err = unmarshalAwsOutput(ret.String(), "TargetGroups", backendgroups)
	if err != nil {
		return nil, err
	}

	if len(backendgroups) == 1 {
		backendgroups[0].region = self
		return &backendgroups[0], nil
	}

	return nil, cloudprovider.ErrNotFound
}