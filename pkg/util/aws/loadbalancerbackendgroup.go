package aws

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/elbv2"
	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SElbBackendGroup struct {
	region *SRegion
	lb     *SElb

	TargetGroupName            string   `json:"TargetGroupName"`
	Protocol                   string   `json:"Protocol"`
	Port                       int64    `json:"Port"`
	VpcID                      string   `json:"VpcId"`
	TargetType                 string   `json:"TargetType"`
	HealthyThresholdCount      int    `json:"HealthyThresholdCount"`
	Matcher                    Matcher  `json:"Matcher"`
	UnhealthyThresholdCount    int    `json:"UnhealthyThresholdCount"`
	HealthCheckPath            string   `json:"HealthCheckPath"`
	HealthCheckProtocol        string   `json:"HealthCheckProtocol"`
	HealthCheckPort            string   `json:"HealthCheckPort"`
	HealthCheckIntervalSeconds int    `json:"HealthCheckIntervalSeconds"`
	HealthCheckTimeoutSeconds  int   `json:"HealthCheckTimeoutSeconds"`
	TargetGroupArn             string   `json:"TargetGroupArn"`
	LoadBalancerArns           []string `json:"LoadBalancerArns"`
}

type Matcher struct {
	HTTPCode string `json:"HttpCode"`
}

func (self *SElbBackendGroup) GetId() string {
	return self.TargetGroupArn
}

func (self *SElbBackendGroup) GetName() string {
	return self.TargetGroupName
}

func (self *SElbBackendGroup) GetGlobalId() string {
	return self.GetId()
}

func (self *SElbBackendGroup) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SElbBackendGroup) Refresh() error {
	panic("implement me")
}

func (self *SElbBackendGroup) IsEmulated() bool {
	return false
}

func (self *SElbBackendGroup) GetMetadata() *jsonutils.JSONDict {
	return jsonutils.NewDict()
}

func (self *SElbBackendGroup) GetProjectId() string {
	return ""
}

func (self *SElbBackendGroup) IsDefault() bool {
	return false
}

func (self *SElbBackendGroup) GetType() string {
	return api.LB_BACKENDGROUP_TYPE_NORMAL
}

func (self *SElbBackendGroup) GetILoadbalancerBackends() ([]cloudprovider.ICloudLoadbalancerBackend, error) {
	backends, err := self.region.GetELbBackends(self.GetId())
	if err != nil {
		return nil, err
	}

	ibackends := make([]cloudprovider.ICloudLoadbalancerBackend, len(backends))
	for i := range backends {
		backends[i].region = self.region
		ibackends[i] = &backends[i]
	}

	return ibackends, nil
}

func (self *SElbBackendGroup) GetILoadbalancerBackendById(backendId string) (cloudprovider.ICloudLoadbalancerBackend, error) {
	backend, err := self.region.GetELbBackend(backendId)
	if err != nil {
		return nil, err
	}

	backend.group = self
	return backend, nil
}

func (self *SElbBackendGroup) GetProtocolType() string {
	switch self.Protocol {
	case "TCP":
		return api.LB_LISTENER_TYPE_TCP
	case "UDP":
		return api.LB_LISTENER_TYPE_UDP
	case "HTTP":
		return api.LB_LISTENER_TYPE_HTTP
	default:
		return ""
	}
}

func (self *SElbBackendGroup) GetScheduler() string {
	return ""
}

func (self *SElbBackendGroup) GetHealthCheck() (*cloudprovider.SLoadbalancerHealthCheck, error) {
	health := &cloudprovider.SLoadbalancerHealthCheck{}
	health.HealthCheckRise = self.HealthyThresholdCount
	health.HealthCheckInterval = self.HealthCheckIntervalSeconds
	health.HealthCheckURI = self.HealthCheckPath
	health.HealthCheckType = self.HealthCheckProtocol
	health.HealthCheckTimeout = self.HealthCheckTimeoutSeconds

	return health, nil
}

func (self *SElbBackendGroup) GetStickySession() (*cloudprovider.SLoadbalancerStickySession, error) {
	return nil, nil
}

func (self *SElbBackendGroup) AddBackendServer(serverId string, weight int, port int) (cloudprovider.ICloudLoadbalancerBackend, error) {
	backend, err := self.region.AddElbBackend(self.GetId(), serverId, weight, port)
	if err != nil {
		return nil, err
	}

	backend.region = self.region
	backend.group = self
	return backend, nil
}

func (self *SElbBackendGroup) RemoveBackendServer(serverId string, weight int, port int) error {
	return self.region.RemoveElbBackend(self.GetId(), serverId, weight, port)
}

func (self *SElbBackendGroup) Delete() error {
	return self.region.DeleteElbBackendGroup(self.GetId())
}

func (self *SElbBackendGroup) Sync(group *cloudprovider.SLoadbalancerBackendGroup) error {
	return self.region.SyncELbBackendGroup(self.GetId(), group)
}

func (self *SRegion) GetELbBackends(backendgroupId string) ([]SElbBackend, error) {
	client, err := self.GetElbV2Client()
	if err != nil {
		return nil, err
	}

	params := &elbv2.DescribeTargetHealthInput{}
	params.SetTargetGroupArn(backendgroupId)
	ret, err := client.DescribeTargetHealth(params)
	if err != nil {
		return nil, err
	}

	backends := []SElbBackend{}
	err = unmarshalAwsOutput(ret.String(), "TargetHealthDescriptions", backends)
	if err != nil {
		return nil, err
	}

	for i := range backends {
		backends[i].region = self
	}

	return backends, nil
}

func (self *SRegion) GetELbBackend(backendId string) (*SElbBackend, error) {
	client, err := self.GetElbV2Client()
	if err != nil {
		return nil, err
	}

	groupId, instanceId, port, err := parseElbBackendId(backendId)
	if err != nil {
		return nil, err
	}

	params := &elbv2.DescribeTargetHealthInput{}
	desc := &elbv2.TargetDescription{}
	desc.SetPort(int64(port))
	desc.SetId(instanceId)
	params.SetTargets([]*elbv2.TargetDescription{desc})
	params.SetTargetGroupArn(groupId)
	ret, err := client.DescribeTargetHealth(params)
	if err != nil {
		return nil, err
	}

	backends := []SElbBackend{}
	err = unmarshalAwsOutput(ret.String(), "TargetHealthDescriptions", backends)
	if err != nil {
		return nil, err
	}

	if len(backends) == 1 {
		backends[0].region = self
		return &backends[0], nil
	}

	return nil, cloudprovider.ErrNotFound
}

func parseElbBackendId(id string) (string, string, int, error) {
	segs := strings.Split(id, "::")
	if len(segs) != 3 {
		return "", "", 0, fmt.Errorf("%s is not a valid backend id", id)
	}

	port, err := strconv.Atoi(segs[2])
	if err != nil {
		return "", "", 0, fmt.Errorf("%s is not a valid backend id, %s", id, err)
	}

	return segs[0], segs[1], port, nil
}

func genElbBackendId(backendgroupId string,serverId string,port int) string {
	return strings.Join([]string{backendgroupId, serverId, strconv.Itoa(port)}, "::")
}

func (self *SRegion) AddElbBackend(backendgroupId, serverId string, weight int, port int) (*SElbBackend, error) {
	client, err := self.GetElbV2Client()
	if err != nil {
		return nil, err
	}

	params := &elbv2.RegisterTargetsInput{}
	params.SetTargetGroupArn(backendgroupId)
	desc := &elbv2.TargetDescription{}
	desc.SetId(serverId)
	desc.SetPort(int64(port))
	params.SetTargets([]*elbv2.TargetDescription{desc})
	ret, err := client.RegisterTargets(params)
	if err != nil {
		return nil, err
	}

	backends := []SElbBackend{}
	err = unmarshalAwsOutput(ret.String(), "TargetHealthDescriptions", backends)
	if err != nil {
		return nil, err
	}

	return self.GetELbBackend(genElbBackendId(backendgroupId, serverId, port))
}

func (self *SRegion) RemoveElbBackend(backendgroupId,serverId string, weight int, port int) error {
	client, err := self.GetElbV2Client()
	if err != nil {
		return err
	}

	params := &elbv2.DeregisterTargetsInput{}
	params.SetTargetGroupArn(backendgroupId)
	desc := &elbv2.TargetDescription{}
	desc.SetId(serverId)
	desc.SetPort(int64(port))
	params.SetTargets([]*elbv2.TargetDescription{desc})
	_, err = client.DeregisterTargets(params)
	if err != nil {
		return err
	}

	return nil
}

func (self *SRegion) DeleteElbBackendGroup(backendgroupId string) error {
	client, err := self.GetElbV2Client()
	if err != nil {
		return err
	}

	params := &elbv2.DeleteTargetGroupInput{}
	params.SetTargetGroupArn(backendgroupId)
	_, err = client.DeleteTargetGroup(params)
	if err != nil {
		return err
	}

	return nil
}

func (self *SRegion) SyncELbBackendGroup(backendgroupId string, group *cloudprovider.SLoadbalancerBackendGroup) error {
	client, err := self.GetElbV2Client()
	if err != nil {
		return err
	}

	params := &elbv2.ModifyTargetGroupInput{}
	params.SetTargetGroupArn(backendgroupId)
	params.SetHealthyThresholdCount(int64(group.HealthCheck.HealthCheckRise))
	params.SetHealthCheckTimeoutSeconds(int64(group.HealthCheck.HealthCheckTimeout))
	params.SetHealthCheckProtocol(group.HealthCheck.HealthCheckType)
	params.SetHealthCheckPath(group.HealthCheck.HealthCheckURI)
	params.SetHealthCheckIntervalSeconds(int64(group.HealthCheck.HealthCheckInterval))
	// params.SetMatcher()
	// params.SetUnhealthyThresholdCount()
	_, err = client.ModifyTargetGroup(params)
	if err != nil {
		return err
	}

	err = self.RemoveElbBackends(backendgroupId)
	if err != nil {
		return err
	}

	return self.AddElbBackends(backendgroupId, group.Backends)
}

func (self *SRegion) RemoveElbBackends(backendgroupId string) error {
	client, err := self.GetElbV2Client()
	if err != nil {
		return err
	}

	params := &elbv2.DeregisterTargetsInput{}
	params.SetTargetGroupArn(backendgroupId)
	_, err = client.DeregisterTargets(params)
	if err != nil {
		return err
	}

	return nil
}

func (self *SRegion) AddElbBackends(backendgroupId string, backends []cloudprovider.SLoadbalancerBackend) error {
	client, err := self.GetElbV2Client()
	if err != nil {
		return err
	}

	params := &elbv2.RegisterTargetsInput{}
	params.SetTargetGroupArn(backendgroupId)
	targets := []*elbv2.TargetDescription{}
	for i := range backends {
		desc := &elbv2.TargetDescription{}
		desc.SetId(backends[i].ID)
		desc.SetPort(int64(backends[i].Port))
		targets = append(targets, desc)
	}

	params.SetTargets(targets)
	_, err = client.RegisterTargets(params)
	if err != nil {
		return err
	}

	return nil
}