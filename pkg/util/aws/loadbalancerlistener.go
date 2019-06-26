package aws

import (
	"strings"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SElbListener struct {
	Port            int             `json:"Port"`
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
	return self.ListenerArn
}

func (self *SElbListener) GetName() string {
	return self.ListenerArn
}

func (self *SElbListener) GetGlobalId() string {
	return self.GetId()
}

func (self *SElbListener) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SElbListener) Refresh() error {
	panic("implement me")
}

func (self *SElbListener) IsEmulated() bool {
	return false
}

func (self *SElbListener) GetMetadata() *jsonutils.JSONDict {
	return jsonutils.NewDict()
}

func (self *SElbListener) GetProjectId() string {
	return ""
}

func (self *SElbListener) GetListenerType() string {
	switch self.Protocol {
	case "TCP":
		return api.LB_LISTENER_TYPE_TCP
	case "UDP":
		return api.LB_LISTENER_TYPE_UDP
	case "HTTP":
		return api.LB_LISTENER_TYPE_HTTP
	case "HTTPS":
		return api.LB_LISTENER_TYPE_HTTPS
	case "TCP_SSL":
		return api.LB_LISTENER_TYPE_TCP
	default:
		return ""
	}
}

func (self *SElbListener) GetListenerPort() int {
	return self.Port
}

func (self *SElbListener) GetScheduler() string {
	return ""
}

func (self *SElbListener) GetAclStatus() string {
	return api.LB_BOOL_OFF
}

func (self *SElbListener) GetAclType() string {
	return ""
}

func (self *SElbListener) GetAclId() string {
	return ""
}

func (self *SElbListener) GetEgressMbps() int {
	return 0
}

func (self *SElbListener) GetHealthCheck() string {
	// todo: implment me
	return ""
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
	strings.ReplaceAll()
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
	// todo: 需要从loadblancer attributes 中获取
	return 0
}

func (self *SElbListener) XForwardedForEnabled() bool {
	return false
}

func (self *SElbListener) GzipEnabled() bool {
	return false
}

func (self *SElbListener) GetCertificateId() string {
	if len(self.Certificates) > 0 {
		return self.Certificates[0].CertificateArn
	}

	return ""
}

func (self *SElbListener) GetTLSCipherPolicy() string {
	panic("implement me")
}

func (self *SElbListener) HTTP2Enabled() bool {
	return false
}

func (self *SElbListener) Start() error {
	return nil
}

func (self *SElbListener) Stop() error {
	return cloudprovider.ErrNotSupported
}

func (self *SElbListener) Sync(listener *cloudprovider.SLoadbalancerListener) error {
	panic("implement me")
}

func (self *SElbListener) Delete() error {
	panic("implement me")
}

func (self *SRegion) GetElbListeners(elbId string) ([]SElbListener, error) {
	return nil, nil
}