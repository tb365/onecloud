package aws

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SElbBackendGroup struct {
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
	panic("implement me")
}

func (self *SElbBackendGroup) GetName() string {
	panic("implement me")
}

func (self *SElbBackendGroup) GetGlobalId() string {
	panic("implement me")
}

func (self *SElbBackendGroup) GetStatus() string {
	panic("implement me")
}

func (self *SElbBackendGroup) Refresh() error {
	panic("implement me")
}

func (self *SElbBackendGroup) IsEmulated() bool {
	panic("implement me")
}

func (self *SElbBackendGroup) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElbBackendGroup) GetProjectId() string {
	panic("implement me")
}

func (self *SElbBackendGroup) IsDefault() bool {
	panic("implement me")
}

func (self *SElbBackendGroup) GetType() string {
	panic("implement me")
}

func (self *SElbBackendGroup) GetILoadbalancerBackends() ([]cloudprovider.ICloudLoadbalancerBackend, error) {
	panic("implement me")
}

func (self *SElbBackendGroup) GetILoadbalancerBackendById(backendId string) (cloudprovider.ICloudLoadbalancerBackend, error) {
	panic("implement me")
}

func (self *SElbBackendGroup) GetProtocolType() string {
	panic("implement me")
}

func (self *SElbBackendGroup) GetScheduler() string {
	panic("implement me")
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

