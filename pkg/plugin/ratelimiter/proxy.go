/**
 * Tencent is pleased to support the open source community by making polaris-go available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package ratelimiter

import (
	"github.com/polarismesh/polaris-go/pkg/model"
	"github.com/polarismesh/polaris-go/pkg/plugin"
	"github.com/polarismesh/polaris-go/pkg/plugin/common"
)

// Proxy proxy of ServiceRateLimiter
type Proxy struct {
	ServiceRateLimiter
	engine model.Engine
}

// SetRealPlugin 设置
func (p *Proxy) SetRealPlugin(plug plugin.Plugin, engine model.Engine) {
	p.ServiceRateLimiter = plug.(ServiceRateLimiter)
	p.engine = engine
}

// InitQuota proxy ServiceRateLimiter InitQuota
func (p *Proxy) InitQuota(criteria *InitCriteria) (QuotaBucket, error) {
	result, err := p.ServiceRateLimiter.InitQuota(criteria)
	return result, err
}

// init 注册proxy
func init() {
	plugin.RegisterPluginProxy(common.TypeRateLimiter, &Proxy{})
}
