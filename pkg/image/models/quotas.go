// Copyright 2019 Yunion
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

package models

import (
	"context"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/tristate"

	"yunion.io/x/onecloud/pkg/cloudcommon/db/quotas"
	"yunion.io/x/onecloud/pkg/image/options"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/util/rbacutils"
)

type SQuotaManager struct {
	quotas.SQuotaBaseManager
}

var QuotaManager *SQuotaManager
var QuotaUsageManager *SQuotaManager

func init() {
	pendingStore := quotas.NewMemoryQuotaStore()

	QuotaUsageManager = &SQuotaManager{
		SQuotaBaseManager: quotas.NewQuotaUsageManager(SQuota{}, "quota_usage_tbl"),
	}

	QuotaManager = &SQuotaManager{
		SQuotaBaseManager: quotas.NewQuotaBaseManager(SQuota{}, "quota_tbl", pendingStore, QuotaUsageManager),
	}
}

type SQuota struct {
	quotas.SQuotaBase

	Image int
}

func (self *SQuota) FetchSystemQuota(scope rbacutils.TRbacScope) {
	base := 0
	if scope == rbacutils.ScopeDomain {
		base = 10
	}
	self.Image = options.Options.DefaultImageQuota * base
}

func (self *SQuota) FetchUsage(ctx context.Context, scope rbacutils.TRbacScope, ownerId mcclient.IIdentityProvider, platform []string) error {
	count := ImageManager.count(scope, ownerId, "", tristate.None, false)
	self.Image = int(count["total"].Count)
	return nil
}

func (self *SQuota) IsEmpty() bool {
	if self.Image > 0 {
		return false
	}
	return true
}

func (self *SQuota) Add(quota quotas.IQuota) {
	squota := quota.(*SQuota)
	self.Image = self.Image + squota.Image
}

func (self *SQuota) Sub(quota quotas.IQuota) {
	squota := quota.(*SQuota)
	self.Image = quotas.NonNegative(self.Image - squota.Image)
}

func (self *SQuota) Update(quota quotas.IQuota) {
	squota := quota.(*SQuota)
	if squota.Image > 0 {
		self.Image = squota.Image
	}
}

func (self *SQuota) Exceed(request quotas.IQuota, quota quotas.IQuota) error {
	err := quotas.NewOutOfQuotaError()
	sreq := request.(*SQuota)
	squota := quota.(*SQuota)
	if sreq.Image > 0 && self.Image > squota.Image {
		err.Add("image", squota.Image, self.Image)
	}
	if err.IsError() {
		return err
	} else {
		return nil
	}
}

func (self *SQuota) ToJSON(prefix string) jsonutils.JSONObject {
	ret := jsonutils.NewDict()
	if self.Image > 0 {
		ret.Add(jsonutils.NewInt(int64(self.Image)), quotas.KeyName(prefix, "image"))
	}
	return ret
}
