package aws

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SElbListener struct {
	Port            int64           `json:"Port"`
	Protocol        string          `json:"Protocol"`
	DefaultActions  []DefaultAction `json:"DefaultActions"`
	SSLPolicy       string          `json:"SslPolicy"`
	Certificates    []Certificate   `json:"Certificates"`
	LoadBalancerArn string          `json:"LoadBalancerArn"`
	ListenerArn     string          `json:"ListenerArn"`
}

type Certificate struct {
	CertificateArn string `json:"CertificateArn"`
}

type DefaultAction struct {
	TargetGroupArn string `json:"TargetGroupArn"`
	Type           string `json:"Type"`
}

func (self *SElbListener) GetId() string {
	panic("implement me")
}

func (self *SElbListener) GetName() string {
	panic("implement me")
}

func (self *SElbListener) GetGlobalId() string {
	panic("implement me")
}

func (self *SElbListener) GetStatus() string {
	panic("implement me")
}

func (self *SElbListener) Refresh() error {
	panic("implement me")
}

func (self *SElbListener) IsEmulated() bool {
	panic("implement me")
}

func (self *SElbListener) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElbListener) GetProjectId() string {
	panic("implement me")
}

func (self *SElbListener) GetListenerType() string {
	panic("implement me")
}

func (self *SElbListener) GetListenerPort() int {
	panic("implement me")
}

func (self *SElbListener) GetScheduler() string {
	panic("implement me")
}

func (self *SElbListener) GetAclStatus() string {
	panic("implement me")
}

func (self *SElbListener) GetAclType() string {
	panic("implement me")
}

func (self *SElbListener) GetAclId() string {
	panic("implement me")
}

func (self *SElbListener) GetEgressMbps() int {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheck() string {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckType() string {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckTimeout() int {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckInterval() int {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckRise() int {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckFail() int {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckReq() string {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckExp() string {
	panic("implement me")
}

func (self *SElbListener) GetBackendGroupId() string {
	panic("implement me")
}

func (self *SElbListener) GetBackendServerPort() int {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckDomain() string {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckURI() string {
	panic("implement me")
}

func (self *SElbListener) GetHealthCheckCode() string {
	panic("implement me")
}

func (self *SElbListener) CreateILoadBalancerListenerRule(rule *cloudprovider.SLoadbalancerListenerRule) (cloudprovider.ICloudLoadbalancerListenerRule, error) {
	panic("implement me")
}

func (self *SElbListener) GetILoadBalancerListenerRuleById(ruleId string) (cloudprovider.ICloudLoadbalancerListenerRule, error) {
	panic("implement me")
}

func (self *SElbListener) GetILoadbalancerListenerRules() ([]cloudprovider.ICloudLoadbalancerListenerRule, error) {
	panic("implement me")
}

func (self *SElbListener) GetStickySession() string {
	panic("implement me")
}

func (self *SElbListener) GetStickySessionType() string {
	panic("implement me")
}

func (self *SElbListener) GetStickySessionCookie() string {
	panic("implement me")
}

func (self *SElbListener) GetStickySessionCookieTimeout() int {
	panic("implement me")
}

func (self *SElbListener) XForwardedForEnabled() bool {
	panic("implement me")
}

func (self *SElbListener) GzipEnabled() bool {
	panic("implement me")
}

func (self *SElbListener) GetCertificateId() string {
	panic("implement me")
}

func (self *SElbListener) GetTLSCipherPolicy() string {
	panic("implement me")
}

func (self *SElbListener) HTTP2Enabled() bool {
	panic("implement me")
}

func (self *SElbListener) Start() error {
	panic("implement me")
}

func (self *SElbListener) Stop() error {
	panic("implement me")
}

func (self *SElbListener) Sync(listener *cloudprovider.SLoadbalancerListener) error {
	panic("implement me")
}

func (self *SElbListener) Delete() error {
	panic("implement me")
}

