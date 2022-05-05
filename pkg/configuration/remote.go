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
	"github.com/polarismesh/polaris-go/pkg/plugin/serverconnector"
)

type remoteConfigFileRepo struct {
	connector          serverconnector.ServerConnector
	configuration      config.Configuration
	configFileMetadata defaultConfigFileMetadata
}

type remoteRepoChangeListener func(configFileMetadata defaultConfigFileMetadata, newContent string)

func newRemoteConfigFileRepo(metadata defaultConfigFileMetadata,
	connector serverconnector.ServerConnector,
	configuration config.Configuration) *remoteConfigFileRepo {
	return &remoteConfigFileRepo{
		connector:          connector,
		configuration:      configuration,
		configFileMetadata: metadata,
	}
}

func (r *remoteConfigFileRepo) getContent() string {
	return ""
}

func (r *remoteConfigFileRepo) addChangeListener(listener remoteRepoChangeListener) {

}
