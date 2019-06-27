package aws

import (
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
	HealthyThresholdCount      int64    `json:"HealthyThresholdCount"`
	Matcher                    Matcher  `json:"Matcher"`
	UnhealthyThresholdCount    int64    `json:"UnhealthyThresholdCount"`
	HealthCheckPath            string   `json:"HealthCheckPath"`
	HealthCheckProtocol        string   `json:"HealthCheckProtocol"`
	HealthCheckPort            string   `json:"HealthCheckPort"`
	HealthCheckIntervalSeconds int64    `json:"HealthCheckIntervalSeconds"`
	HealthCheckTimeoutSeconds  int64    `json:"HealthCheckTimeoutSeconds"`
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
	panic("implement me")
}

func (self *SElbBackendGroup) GetILoadbalancerBackendById(backendId string) (cloudprovider.ICloudLoadbalancerBackend, error) {
	panic("implement me")
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
	panic("implement me")
}

func (self *SElbBackendGroup) GetStickySession() (*cloudprovider.SLoadbalancerStickySession, error) {
	panic("implement me")
}

func (self *SElbBackendGroup) AddBackendServer(serverId string, weight int, port int) (cloudprovider.ICloudLoadbalancerBackend, error) {
	panic("implement me")
}

func (self *SElbBackendGroup) RemoveBackendServer(serverId string, weight int, port int) error {
	panic("implement me")
}

func (self *SElbBackendGroup) Delete() error {
	panic("implement me")
}

func (self *SElbBackendGroup) Sync(group *cloudprovider.SLoadbalancerBackendGroup) error {
	panic("implement me")
}

