package aws

import (
	"time"

	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
)

type SElbCertificate struct {
	region *SRegion

	CertificateArn string `json:"CertificateArn"`
	IsDefault      bool   `json:"IsDefault"`
}

func (self *SElbCertificate) GetId() string {
	return self.CertificateArn
}

func (self *SElbCertificate) GetName() string {
	return self.GetId()
}

func (self *SElbCertificate) GetGlobalId() string {
	return self.GetId()
}

func (self *SElbCertificate) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SElbCertificate) Refresh() error {
	panic("implement me")
}

func (self *SElbCertificate) IsEmulated() bool {
	return false
}

func (self *SElbCertificate) GetMetadata() *jsonutils.JSONDict {
	return jsonutils.NewDict()
}

func (self *SElbCertificate) GetProjectId() string {
	return ""
}

func (self *SElbCertificate) Sync(name, privateKey, publickKey string) error {
	panic("implement me")
}

func (self *SElbCertificate) Delete() error {
	panic("implement me")
}

func (self *SElbCertificate) GetCommonName() string {
	panic("implement me")
}

func (self *SElbCertificate) GetSubjectAlternativeNames() string {
	panic("implement me")
}

func (self *SElbCertificate) GetFingerprint() string {
	panic("implement me")
}

func (self *SElbCertificate) GetExpireTime() time.Time {
	panic("implement me")
}

func (self *SElbCertificate) GetPublickKey() string {
	panic("implement me")
}

func (self *SElbCertificate) GetPrivateKey() string {
	panic("implement me")
}
