// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeAlarmRequest struct {

    core.JDCloudRequest

    /* 规则id  */
    AlarmId string `json:"alarmId"`
}

/*
 * param alarmId: 规则id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmRequest(
    alarmId string,
) *DescribeAlarmRequest {

	return &DescribeAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/groupAlarms/{alarmId}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        AlarmId: alarmId,
	}
}

/*
 * param alarmId: 规则id (Required)
 */
func NewDescribeAlarmRequestWithAllParams(
    alarmId string,
) *DescribeAlarmRequest {

    return &DescribeAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarms/{alarmId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        AlarmId: alarmId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmRequestWithoutParam() *DescribeAlarmRequest {

    return &DescribeAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/groupAlarms/{alarmId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param alarmId: 规则id(Required) */
func (r *DescribeAlarmRequest) SetAlarmId(alarmId string) {
    r.AlarmId = alarmId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmRequest) GetRegionId() string {
    return ""
}

type DescribeAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmResult `json:"result"`
}

type DescribeAlarmResult struct {
    AlarmId string `json:"alarmId"`
    AlarmStatus int64 `json:"alarmStatus"`
    AlarmStatusList []int64 `json:"alarmStatusList"`
    BaseContact []monitor.BaseContact `json:"baseContact"`
    CreateTime string `json:"createTime"`
    Dimension string `json:"dimension"`
    DimensionName string `json:"dimensionName"`
    Enabled int64 `json:"enabled"`
    MultiWebHook []monitor.WebHookOption `json:"multiWebHook"`
    NoticeOption []monitor.NoticeOption `json:"noticeOption"`
    Product string `json:"product"`
    ProductName string `json:"productName"`
    ResourceOption monitor.ResourceOption `json:"resourceOption"`
    RuleName string `json:"ruleName"`
    RuleOption monitor.RuleOptionDetail `json:"ruleOption"`
    RuleType string `json:"ruleType"`
    RuleVersion string `json:"ruleVersion"`
    Tags interface{} `json:"tags"`
    WebHookOption monitor.WebHookOption `json:"webHookOption"`
}