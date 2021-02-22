package monitor

import (
	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/mcclient/options"
)

type AlertRecordListOptions struct {
	options.BaseListOptions

	AlertId  string   `help:"id of alert"`
	Level    string   `help:"alert level"`
	State    string   `help:"alert state"`
	ResTypes []string `json:"res_types"`
	Alerting bool     `json:"alerting"`
}

func (o *AlertRecordListOptions) Params() (jsonutils.JSONObject, error) {
	return options.ListStructToParams(o)
}

type AlertRecordShowOptions struct {
	ID string `help:"ID of Metric " json:"-"`
}

func (o *AlertRecordShowOptions) Params() (jsonutils.JSONObject, error) {
	return options.StructToParams(o)
}

func (o *AlertRecordShowOptions) GetId() string {
	return o.ID
}

type AlertRecordTotalOptions struct {
	ID string `help:"total-alert" json:"-"`
	options.BaseListOptions
}

func (o *AlertRecordTotalOptions) Params() (jsonutils.JSONObject, error) {
	return options.ListStructToParams(o)
}

func (o *AlertRecordTotalOptions) GetId() string {
	return o.ID
}
