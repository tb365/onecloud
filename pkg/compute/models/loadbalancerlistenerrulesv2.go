package models

import (
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
			SLoadbalancerListenerRule{},
			"loadbalancerlistenerrulesv2_tbl",
			"loadbalancerlistenerrulev2",
			"loadbalancerlistenerrulesv2",
		),
	}
	LoadbalancerListenerRuleManager.SetVirtualObject(LoadbalancerListenerRuleManager)
}

/*
SLoadbalancerListenerRuleV2 不兼容SLoadbalancerListenerRule。
相比于SLoadbalancerListenerRule表，去掉了domian 和path字段。新增condition字段。通过特定格式的JSON字符串，
支持多个条件(eg. Host is baidu.com OR *.google.com && Path is /testOR /test2 OR /test3 )的组合。

*/
type SLoadbalancerListenerRuleV2 struct {
	db.SVirtualResourceBase
	db.SExternalizedResourceBase

	SManagedResourceBase
	SCloudregionResourceBase

	ListenerId     string `width:"36" charset:"ascii" nullable:"false" list:"user" create:"optional"`
	BackendGroupId string `width:"36" charset:"ascii" nullable:"false" list:"user" create:"optional" update:"user"`

	Condition string `charset:"ascii" nullable:"false" list:"user" create:"optional"`

	SLoadbalancerHealthCheck // 目前只有腾讯云HTTP、HTTPS类型的健康检查是和规则绑定的。
	SLoadbalancerHTTPRateLimiter
}
