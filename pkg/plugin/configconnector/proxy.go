/*
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
 *  under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 */

package configconnector

import (
	"github.com/polarismesh/polaris-go/pkg/model"
	configpb "github.com/polarismesh/polaris-go/pkg/model/pb/v1"
	"github.com/polarismesh/polaris-go/pkg/plugin"
	"github.com/polarismesh/polaris-go/pkg/plugin/common"
)

type Proxy struct {
	ConfigConnector
	engine model.Engine
}

func (p *Proxy) SetRealPlugin(plugin plugin.Plugin, engine model.Engine) {
	p.ConfigConnector = plugin.(ConfigConnector)
	p.engine = engine
}

// GetConfigFile Get config file
func (p *Proxy) GetConfigFile(namespace, group, filename string) (*configpb.ConfigClientResponse, error) {
	response, err := p.ConfigConnector.GetConfigFile(namespace, group, filename)
	return response, err
}

// WatchConfigFiles Watch config files
func (p *Proxy) WatchConfigFiles(request *configpb.ClientWatchConfigFileRequest) (*configpb.ConfigClientResponse, error) {
	response, err := p.ConfigConnector.WatchConfigFiles(request)
	return response, err
}

// init 注册proxy
func init() {
	plugin.RegisterPluginProxy(common.TypeConfigConnector, &Proxy{})
}
