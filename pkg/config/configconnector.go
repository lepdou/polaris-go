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

package config

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"

	"github.com/polarismesh/polaris-go/pkg/plugin/common"
)

// ConfigConnectorConfigImpl 对接配置中心连接器相关配置
type ConfigConnectorConfigImpl struct {
	ServerConnectorConfigImpl

	ConnectorType string `yaml:"connectorType" json:"connectorType"`
}

// GetConnectorType 获取连接器类型
func (c *ConfigConnectorConfigImpl) GetConnectorType() string {
	return c.ConnectorType
}

// SetConnectorType 设置连接器类型
func (c *ConfigConnectorConfigImpl) SetConnectorType(connectorType string) {
	c.ConnectorType = connectorType
}

// Verify 检验ConfigConnector配置
func (c *ConfigConnectorConfigImpl) Verify() error {
	if nil == c {
		return errors.New("ConfigConnectorConfig is nil")
	}
	var errs error
	var err error
	if err = c.ServerConnectorConfigImpl.Verify(); err != nil {
		errs = multierror.Append(errs, err)
	}
	if len(c.ConnectorType) == 0 {
		errs = multierror.Append(errs, fmt.Errorf("config.serverConnector.connectorType is empty"))
	}
	return errs
}

// SetDefault 设置ConfigConnector配置的默认值
func (c *ConfigConnectorConfigImpl) SetDefault() {
	c.ServerConnectorConfigImpl.SetDefaultWithoutType()
	if len(c.Addresses) == 0 {
		c.SetAddresses([]string{DefaultConfigConnectorAddresses})
	}
	if len(c.ConnectorType) == 0 {
		c.ConnectorType = DefaultConnectorType
	}
	c.Plugin.SetDefault(common.TypeConfigConnector)
}

// Init 配置初始化
func (c *ConfigConnectorConfigImpl) Init() {
	c.Plugin = PluginConfigs{}
	c.Plugin.Init(common.TypeConfigConnector)
}
