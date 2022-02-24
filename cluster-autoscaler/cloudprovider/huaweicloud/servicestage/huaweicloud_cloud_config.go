/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package huaweicloud

import (
	huaweicloudsdkbasic "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	huaweicloudsdkconfig "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	huaweicloudsdkcce "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3"
	huaweicloudsdkecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
)

// CloudConfig is the cloud config file for huaweicloud.
type CloudConfig struct {
	Global struct {
		CCEEndpoint string `gcfg:"cce-endpoint"`
		ECSEndpoint string `gcfg:"ecs-endpoint"`
		ProjectID   string `gcfg:"project-id"`
		AccessKey   string `gcfg:"access-key"`
		SecretKey   string `gcfg:"secret-key"`
		Region      string `gcfg:"region"`    // example: "cn-north-4"
		DomainID    string `gcfg:"domain-id"` // The ACCOUNT ID. example: "a0e8ff63c0fb4fd49cc2dbdf1dea14e2"
	}
}

func (c *CloudConfig) getCCEClient() *huaweicloudsdkcce.CceClient {
	// There are two types of services provided by HUAWEI CLOUD according to scope:
	// - Regional services: most of services belong to this classification, such as ECS.
	// - Global services: such as IAM, TMS, EPS.
	// For Regional services' authentication, projectId is required.
	// For global services' authentication, domainId is required.
	// More details please refer to:
	// https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/0281b9734f0f95ed5565729e54d96e9820262426/README.md#use-go-sdk
	credentials := huaweicloudsdkbasic.NewCredentialsBuilder().
		WithAk(c.Global.AccessKey).
		WithSk(c.Global.SecretKey).
		WithProjectId(c.Global.ProjectID).
		Build()

	client := huaweicloudsdkcce.CceClientBuilder().
		WithEndpoint(c.Global.CCEEndpoint).
		WithCredential(credentials).
		WithHttpConfig(huaweicloudsdkconfig.DefaultHttpConfig()).
		Build()

	return huaweicloudsdkcce.NewCceClient(client)
}

func (c *CloudConfig) getECSClient() *huaweicloudsdkecs.EcsClient {
	// There are two types of services provided by HUAWEI CLOUD according to scope:
	// - Regional services: most of services belong to this classification, such as ECS.
	// - Global services: such as IAM, TMS, EPS.
	// For Regional services' authentication, projectId is required.
	// For global services' authentication, domainId is required.
	// More details please refer to:
	// https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/0281b9734f0f95ed5565729e54d96e9820262426/README.md#use-go-sdk
	credentials := huaweicloudsdkbasic.NewCredentialsBuilder().
		WithAk(c.Global.AccessKey).
		WithSk(c.Global.SecretKey).
		WithProjectId(c.Global.ProjectID).
		Build()

	client := huaweicloudsdkecs.EcsClientBuilder().
		WithEndpoint(c.Global.ECSEndpoint).
		WithCredential(credentials).
		WithHttpConfig(huaweicloudsdkconfig.DefaultHttpConfig()).
		Build()

	return huaweicloudsdkecs.NewEcsClient(client)
}
