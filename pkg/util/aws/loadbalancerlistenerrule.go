package aws

import (
	"yunion.io/x/jsonutils"
	api "yunion.io/x/onecloud/pkg/apis/compute"
)

type SElbListenerRule struct {
	Priority   string      `json:"Priority"`
	IsDefault  bool        `json:"IsDefault"`
	Actions    []Action    `json:"Actions"`
	RuleArn    string      `json:"RuleArn"`
	Conditions []Condition `json:"Conditions"`
}

type Action struct {
	TargetGroupArn string `json:"TargetGroupArn"`
	Type           string `json:"Type"`
}

type Condition struct {
	Field  string   `json:"Field"`
	Values []string `json:"Values"`
}


func (self *SElbListenerRule) GetId() string {
	return self.RuleArn
}

func (self *SElbListenerRule) GetName() string {
	return self.RuleArn
}

func (self *SElbListenerRule) GetGlobalId() string {
	return self.GetId()
}

func (self *SElbListenerRule) GetStatus() string {
	return api.LB_STATUS_ENABLED
}

func (self *SElbListenerRule) Refresh() error {
	panic("implement me")
}

func (self *SElbListenerRule) IsEmulated() bool {
	panic("implement me")
}

func (self *SElbListenerRule) GetMetadata() *jsonutils.JSONDict {
	panic("implement me")
}

func (self *SElbListenerRule) GetProjectId() string {
	panic("implement me")
}

func (self *SElbListenerRule) GetDomain() string {
	panic("implement me")
}

func (self *SElbListenerRule) GetPath() string {
	panic("implement me")
}

func (self *SElbListenerRule) GetBackendGroupId() string {
	panic("implement me")
}

func (self *SElbListenerRule) Delete() error {
	panic("implement me")
}

