package aws

import (
	"time"
	"yunion.io/x/jsonutils"
)

type SElbCertificate struct {

}

func (self *SElbCertificate) GetId() string {
	panic("implement me")
}

func (self *SElbCertificate) GetName() string {
	panic("implement me")
}

func (self *SElbCertificate) GetGlobalId() string {
	panic("implement me")
}

func (self *SElbCertificate) GetStatus() string {
	panic("implement me")
}

func (self *SElbCertificate) Refresh() error {
	panic("implement me")
}

func (self *SElbCertificate) IsEmulated() bool {
	panic("implement me")
}

func (self *SElbCertificate) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElbCertificate) GetProjectId() string {
	panic("implement me")
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

