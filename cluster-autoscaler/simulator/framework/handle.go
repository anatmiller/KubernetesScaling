/*
Copyright 2024 The Kubernetes Authors.

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

package framework

import (
	"context"
	"fmt"

	"k8s.io/client-go/informers"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
	scheduler_config "k8s.io/kubernetes/pkg/scheduler/apis/config/latest"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
	scheduler_plugins "k8s.io/kubernetes/pkg/scheduler/framework/plugins"
	draplugin "k8s.io/kubernetes/pkg/scheduler/framework/plugins/dynamicresources"
	schedulerframeworkruntime "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
	schedulermetrics "k8s.io/kubernetes/pkg/scheduler/metrics"
)

// Handle is meant for interacting with the scheduler framework.
type Handle struct {
	Framework        schedulerframework.Framework
	DelegatingLister *DelegatingSchedulerSharedLister
}

// NewHandle builds a framework Handle based on the provided informers and scheduler config.
func NewHandle(informerFactory informers.SharedInformerFactory, schedConfig *config.KubeSchedulerConfiguration, draEnabled bool) (*Handle, error) {
	if schedConfig == nil {
		var err error
		schedConfig, err = scheduler_config.Default()
		if err != nil {
			return nil, fmt.Errorf("couldn't create scheduler config: %v", err)
		}
	}
	if len(schedConfig.Profiles) != 1 {
		return nil, fmt.Errorf("unexpected scheduler config: expected one scheduler profile only (found %d profiles)", len(schedConfig.Profiles))
	}
	schedProfile := &schedConfig.Profiles[0]

	sharedLister := NewDelegatingSchedulerSharedLister()
	opts := []schedulerframeworkruntime.Option{
		schedulerframeworkruntime.WithInformerFactory(informerFactory),
		schedulerframeworkruntime.WithSnapshotSharedLister(sharedLister),
	}

	if draEnabled {
		schedProfile = profileWithDraPlugin(schedProfile)
		opts = append(opts, schedulerframeworkruntime.WithSharedDRAManager(sharedLister))
	}

	schedulermetrics.InitMetrics()
	framework, err := schedulerframeworkruntime.NewFramework(
		context.TODO(),
		scheduler_plugins.NewInTreeRegistry(),
		schedProfile,
		opts...,
	)

	if err != nil {
		return nil, fmt.Errorf("couldn't create scheduler framework; %v", err)
	}

	return &Handle{
		Framework:        framework,
		DelegatingLister: sharedLister,
	}, nil
}

func profileWithDraPlugin(profile *config.KubeSchedulerProfile) *config.KubeSchedulerProfile {
	result := profile.DeepCopy()
	addPluginIfNotPresent(result.Plugins.PreFilter, draplugin.Name)
	addPluginIfNotPresent(result.Plugins.Filter, draplugin.Name)
	addPluginIfNotPresent(result.Plugins.Reserve, draplugin.Name)
	return result
}

func addPluginIfNotPresent(pluginSet config.PluginSet, pluginName string) {
	for _, plugin := range pluginSet.Enabled {
		if plugin.Name == pluginName {
			// Plugin already present in the set.
			return
		}
	}
	// Plugin not present in the set, add it.
	pluginSet.Enabled = append(pluginSet.Enabled, config.Plugin{Name: pluginName})
}
