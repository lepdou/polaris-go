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
	"github.com/polarismesh/polaris-go/pkg/model"
)

const (
	initVersion = 0
)

type defaultConfigFile struct {
	model.DefaultConfigFileMetadata
	remoteConfigFileRepo *remoteConfigFileRepo
}

func newDefaultConfigFile(metadata model.DefaultConfigFileMetadata, repo *remoteConfigFileRepo) *defaultConfigFile {
	configFile := &defaultConfigFile{
		remoteConfigFileRepo: repo,
	}
	configFile.Namespace = metadata.GetNamespace()
	configFile.FileGroup = metadata.GetFileGroup()
	configFile.FileName = metadata.GetFileName()

	return configFile
}

// GetContent 获取配置文件内容
func (c *defaultConfigFile) GetContent() string {
	return ""
}

// HasContent 是否有配置内容
func (c *defaultConfigFile) HasContent() bool {
	return false
}

// AddChangeListenerWithChannel 增加配置文件变更监听器，返回 chan 对象
func (c *defaultConfigFile) AddChangeListenerWithChannel() chan model.ConfigFileChangeEvent {
	return nil
}

// AddChangeListener 增加配置文件变更监听器
func (c *defaultConfigFile) AddChangeListener(cb model.OnConfigFileChange) {

}
