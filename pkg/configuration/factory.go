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

package configuration

import (
	"github.com/polarismesh/polaris-go/pkg/config"
	"github.com/polarismesh/polaris-go/pkg/model"
	"github.com/polarismesh/polaris-go/pkg/plugin/serverconnector"
	"sync"
)

type configFileFactory struct {
	connector     serverconnector.ServerConnector
	configuration config.Configuration
}

func newConfigFileFactory(connector serverconnector.ServerConnector, configuration config.Configuration) *configFileFactory {
	return &configFileFactory{
		connector:     connector,
		configuration: configuration,
	}
}

func (c *configFileFactory) createConfigFile(configFileMetadata model.DefaultConfigFileMetadata) model.ConfigFile {
	configFileRemoteRepo := newRemoteConfigFileRepo(configFileMetadata, c.connector, c.configuration)
	return newDefaultConfigFile(configFileMetadata, configFileRemoteRepo)
}

type configFileFactoryManager struct {
	factories     *sync.Map
	connector     serverconnector.ServerConnector
	configuration config.Configuration
}

func newConfigFileFactoryManager(connector serverconnector.ServerConnector, configuration config.Configuration) *configFileFactoryManager {
	return &configFileFactoryManager{
		factories:     new(sync.Map),
		connector:     connector,
		configuration: configuration,
	}
}

func (c *configFileFactoryManager) getFactory(configFileMetadata model.DefaultConfigFileMetadata) *configFileFactory {
	factoryObj, ok := c.factories.Load(configFileMetadata)
	if ok {
		return factoryObj.(*configFileFactory)
	}

	factory := newConfigFileFactory(c.connector, c.configuration)

	c.factories.Store(configFileMetadata, factory)

	return factory
}
