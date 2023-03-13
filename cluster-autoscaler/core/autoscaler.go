/*
Copyright 2016 The Kubernetes Authors.

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

package core

import (
	"strings"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	cloudBuilder "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/builder"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/autoscaler/cluster-autoscaler/core/scaledown/pdb"
	"k8s.io/autoscaler/cluster-autoscaler/core/scaleup"
	"k8s.io/autoscaler/cluster-autoscaler/debuggingsnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/estimator"
	"k8s.io/autoscaler/cluster-autoscaler/expander"
	"k8s.io/autoscaler/cluster-autoscaler/expander/factory"
	ca_processors "k8s.io/autoscaler/cluster-autoscaler/processors"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/predicatechecker"
	"k8s.io/autoscaler/cluster-autoscaler/utils/backoff"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	kube_client "k8s.io/client-go/kubernetes"
)

// AutoscalerOptions is the whole set of options for configuring an autoscaler
type AutoscalerOptions struct {
	config.AutoscalingOptions
	KubeClient             kube_client.Interface
	EventsKubeClient       kube_client.Interface
	AutoscalingKubeClients *context.AutoscalingKubeClients
	CloudProvider          cloudprovider.CloudProvider
	PredicateChecker       predicatechecker.PredicateChecker
	ClusterSnapshot        clustersnapshot.ClusterSnapshot
	ExpanderStrategy       expander.Strategy
	EstimatorBuilder       estimator.EstimatorBuilder
	Processors             *ca_processors.AutoscalingProcessors
	Backoff                backoff.Backoff
	DebuggingSnapshotter   debuggingsnapshot.DebuggingSnapshotter
	RemainingPdbTracker    pdb.RemainingPdbTracker
	ScaleUpManagerFactory  scaleup.ManagerFactory
}

// Autoscaler is the main component of CA which scales up/down node groups according to its configuration
// The configuration can be injected at the creation of an autoscaler
type Autoscaler interface {
	// Start starts components running in background.
	Start() error
	// RunOnce represents an iteration in the control-loop of CA
	RunOnce(currentTime time.Time) errors.AutoscalerError
	// ExitCleanUp is a clean-up performed just before process termination.
	ExitCleanUp()
}

// NewAutoscaler creates an autoscaler of an appropriate type according to the parameters
func NewAutoscaler(opts AutoscalerOptions) (Autoscaler, errors.AutoscalerError) {
	err := initializeDefaultOptions(&opts)
	if err != nil {
		return nil, errors.ToAutoscalerError(errors.InternalError, err)
	}
	return NewStaticAutoscaler(
		opts.AutoscalingOptions,
		opts.PredicateChecker,
		opts.ClusterSnapshot,
		opts.AutoscalingKubeClients,
		opts.Processors,
		opts.CloudProvider,
		opts.ExpanderStrategy,
		opts.EstimatorBuilder,
		opts.Backoff,
		opts.DebuggingSnapshotter,
		opts.RemainingPdbTracker,
		opts.ScaleUpManagerFactory), nil
}

// Initialize default options if not provided.
func initializeDefaultOptions(opts *AutoscalerOptions) error {
	if opts.Processors == nil {
		opts.Processors = ca_processors.DefaultProcessors()
	}
	if opts.AutoscalingKubeClients == nil {
		opts.AutoscalingKubeClients = context.NewAutoscalingKubeClients(opts.AutoscalingOptions, opts.KubeClient, opts.EventsKubeClient)
	}
	if opts.ClusterSnapshot == nil {
		opts.ClusterSnapshot = clustersnapshot.NewBasicClusterSnapshot()
	}
	if opts.RemainingPdbTracker == nil {
		opts.RemainingPdbTracker = pdb.NewBasicRemainingPdbTracker()
	}
	if opts.CloudProvider == nil {
		opts.CloudProvider = cloudBuilder.NewCloudProvider(opts.AutoscalingOptions)
	}
	if opts.ExpanderStrategy == nil {
		expanderFactory := factory.NewFactory()
		expanderFactory.RegisterDefaultExpanders(opts.CloudProvider, opts.AutoscalingKubeClients, opts.KubeClient, opts.ConfigNamespace, opts.GRPCExpanderCert, opts.GRPCExpanderURL)
		expanderStrategy, err := expanderFactory.Build(strings.Split(opts.ExpanderNames, ","))
		if err != nil {
			return err
		}
		opts.ExpanderStrategy = expanderStrategy
	}
	if opts.EstimatorBuilder == nil {
		estimatorBuilder, err := estimator.NewEstimatorBuilder(opts.EstimatorName, estimator.NewThresholdBasedEstimationLimiter(opts.MaxNodesPerScaleUp, opts.MaxNodeGroupBinpackingDuration))
		if err != nil {
			return err
		}
		opts.EstimatorBuilder = estimatorBuilder
	}
	if opts.Backoff == nil {
		opts.Backoff =
			backoff.NewIdBasedExponentialBackoff(opts.InitialNodeGroupBackoffDuration, opts.MaxNodeGroupBackoffDuration, opts.NodeGroupBackoffResetTimeout)
	}

	return nil
}
