package models

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
)

type SLoadbalancerListenerRuleV2Manager struct {
	SLoadbalancerLogSkipper
	db.SVirtualResourceBaseManager
}

var LoadbalancerListenerRuleV2Manager *SLoadbalancerListenerRuleV2Manager

func init() {
	LoadbalancerListenerRuleV2Manager = &SLoadbalancerListenerRuleV2Manager{
		SVirtualResourceBaseManager: db.NewVirtualResourceBaseManager(
			SLoadbalancerListenerRuleV2{},
			"loadbalancerlistenerrulesv2_tbl",
			"loadbalancerlistenerrulev2",
			"loadbalancerlistenerrulesv2",
		),
	}
	LoadbalancerListenerRuleV2Manager.SetVirtualObject(LoadbalancerListenerRuleV2Manager)
}

/*
date 2019.06.25:
相比LoadbalancerListenerRule，新增condition字段。通过特定格式的JSON字符串，
支持多个条件(eg. Host is baidu.com OR *.google.com && Path is /testOR /test2 OR /test3 )的组合。
新接入elb不再使用domian 和path字段。原有已接入elb保持不变，后续逐步，统一到LoadbalancerListenerRuleV2中。
*/
type SLoadbalancerListenerRuleV2 struct {
	db.SVirtualResourceBase
	db.SExternalizedResourceBase

	SManagedResourceBase
	SCloudregionResourceBase

	ListenerId     string `width:"36" charset:"ascii" nullable:"false" list:"user" create:"optional"`
	BackendGroupId string `width:"36" charset:"ascii" nullable:"false" list:"user" create:"optional" update:"user"`

	Condition string `charset:"ascii" nullable:"false" list:"user" create:"optional"`
}

func ListenerRuleParser(condition string) (jsonutils.JSONObject, error) {

}