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

type configFileManager struct {
	configFileCache          *sync.Map
	lock                     *sync.Mutex
	connector                serverconnector.ServerConnector
	configuration            config.Configuration
	configFileFactoryManager *configFileFactoryManager
}

func newConfigFileManager(connector serverconnector.ServerConnector,
	configuration config.Configuration) *configFileManager {
	return &configFileManager{
		configFileCache:          new(sync.Map),
		lock:                     new(sync.Mutex),
		connector:                connector,
		configuration:            configuration,
		configFileFactoryManager: newConfigFileFactoryManager(connector, configuration),
	}
}

func (c *configFileManager) getConfigFile(configFileMetadata model.DefaultConfigFileMetadata) (model.ConfigFile, error) {
	configFileObj, ok := c.configFileCache.Load(configFileMetadata)
	if !ok {
		c.lock.Lock()
		defer c.lock.Unlock()

		configFileObj, ok = c.configFileCache.Load(configFileMetadata)
		if !ok {
			configFileFactory := c.configFileFactoryManager.getFactory(configFileMetadata)

			configFile := configFileFactory.createConfigFile(configFileMetadata)

			c.configFileCache.Store(configFileMetadata, configFile)

			return configFile, nil
		} else {
			return configFileObj.(model.ConfigFile), nil
		}
	} else {
		return configFileObj.(model.ConfigFile), nil
	}
}
