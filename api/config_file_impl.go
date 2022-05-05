/**
 * Tencent is pleased to support the open source community by making Polaris available.
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

package api

import (
	"github.com/polarismesh/polaris-go/pkg/config"
	"github.com/polarismesh/polaris-go/pkg/model"
)

type configFileAPI struct {
	context SDKContext
}

func newConfigFileAPI() (ConfigFileAPI, error) {
	return newConfigFileAPIByConfig(config.NewDefaultConfigurationWithDomain())
}

// newConfigFileAPIByConfig 通过配置对象创建SDK ConfigFileAPI 对象
func newConfigFileAPIByConfig(cfg config.Configuration) (ConfigFileAPI, error) {
	context, err := InitContextByConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &configFileAPI{context}, nil
}

func newConfigFileAPIBySDKContext(context SDKContext) ConfigFileAPI {
	return &configFileAPI{
		context: context,
	}
}

// newConsumerAPIByAddress 通过address创建ConsumerAPI
func newConfigFileAPIByAddress(address ...string) (ConfigFileAPI, error) {
	conf := config.NewDefaultConfiguration(address)
	return newConfigFileAPIByConfig(conf)
}

func (c *configFileAPI) GetConfigFile(namespace, fileGroup, fileName string) (model.ConfigFile, error) {
	return c.context.GetEngine().SyncGetConfigFile(namespace, fileGroup, fileName)
}

func (c *configFileAPI) SDKContext() SDKContext {
	return c.context
}
