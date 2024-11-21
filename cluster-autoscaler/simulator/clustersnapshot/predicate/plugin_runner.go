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

package predicate

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/framework"

	apiv1 "k8s.io/api/core/v1"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

// SchedulerPluginRunner can be used to run various phases of scheduler plugins through the scheduler framework.
type SchedulerPluginRunner struct {
	fwHandle  *framework.Handle
	snapshot  clustersnapshot.ClusterSnapshot
	lastIndex int
}

// NewSchedulerPluginRunner builds a SchedulerPluginRunner.
func NewSchedulerPluginRunner(fwHandle *framework.Handle, snapshot clustersnapshot.ClusterSnapshot) *SchedulerPluginRunner {
	return &SchedulerPluginRunner{fwHandle: fwHandle, snapshot: snapshot}
}

// RunFiltersUntilPassingNode runs the scheduler framework PreFilter phase once, and then keeps running the Filter phase for all nodes in the cluster that match the provided
// function - until a Node where the filters pass is found. Filters are only run for matching Nodes. If no matching node with passing filters is found, an error is returned.
//
// The node iteration always starts from the next Node from the last Node that was found by this method. TODO: Extract the iteration strategy out of SchedulerPluginRunner.
func (p *SchedulerPluginRunner) RunFiltersUntilPassingNode(pod *apiv1.Pod, nodeMatches func(*framework.NodeInfo) bool) (*apiv1.Node, *schedulerframework.CycleState, clustersnapshot.SchedulingError) {
	nodeInfosList, err := p.snapshot.ListNodeInfos()
	if err != nil {
		return nil, nil, clustersnapshot.NewSchedulingInternalError(pod, "ClusterSnapshot not provided")
	}

	p.fwHandle.DelegatingLister.UpdateDelegate(p.snapshot)
	defer p.fwHandle.DelegatingLister.ResetDelegate()

	state := schedulerframework.NewCycleState()
	preFilterResult, preFilterStatus, _ := p.fwHandle.Framework.RunPreFilterPlugins(context.TODO(), state, pod)
	if !preFilterStatus.IsSuccess() {
		return nil, nil, clustersnapshot.NewFailingPredicateError(pod, preFilterStatus.Plugin(), preFilterStatus.Reasons(), "PreFilter failed", "")
	}

	for i := range nodeInfosList {
		nodeInfo := nodeInfosList[(p.lastIndex+i)%len(nodeInfosList)]
		if !nodeMatches(nodeInfo) {
			continue
		}

		if !preFilterResult.AllNodes() && !preFilterResult.NodeNames.Has(nodeInfo.Node().Name) {
			continue
		}

		// Be sure that the node is schedulable.
		if nodeInfo.Node().Spec.Unschedulable {
			continue
		}

		filterStatus := p.fwHandle.Framework.RunFilterPlugins(context.TODO(), state, pod, nodeInfo.ToScheduler())
		if filterStatus.IsSuccess() {
			p.lastIndex = (p.lastIndex + i + 1) % len(nodeInfosList)
			return nodeInfo.Node(), state, nil
		}
	}
	return nil, nil, clustersnapshot.NewNoNodesPassingPredicatesFoundError(pod)
}

// RunFiltersOnNode runs the scheduler framework PreFilter and Filter phases to check if the given pod can be scheduled on the given node.
func (p *SchedulerPluginRunner) RunFiltersOnNode(pod *apiv1.Pod, nodeName string) (*apiv1.Node, *schedulerframework.CycleState, clustersnapshot.SchedulingError) {
	nodeInfo, err := p.snapshot.GetNodeInfo(nodeName)
	if err != nil {
		return nil, nil, clustersnapshot.NewSchedulingInternalError(pod, fmt.Sprintf("error obtaining NodeInfo for name %q: %v", nodeName, err))
	}

	p.fwHandle.DelegatingLister.UpdateDelegate(p.snapshot)
	defer p.fwHandle.DelegatingLister.ResetDelegate()

	state := schedulerframework.NewCycleState()
	_, preFilterStatus, _ := p.fwHandle.Framework.RunPreFilterPlugins(context.TODO(), state, pod)
	if !preFilterStatus.IsSuccess() {
		return nil, nil, clustersnapshot.NewFailingPredicateError(pod, preFilterStatus.Plugin(), preFilterStatus.Reasons(), "PreFilter failed", "")
	}

	filterStatus := p.fwHandle.Framework.RunFilterPlugins(context.TODO(), state, pod, nodeInfo.ToScheduler())

	if !filterStatus.IsSuccess() {
		filterName := filterStatus.Plugin()
		filterReasons := filterStatus.Reasons()
		unexpectedErrMsg := ""
		if !filterStatus.IsRejected() {
			unexpectedErrMsg = fmt.Sprintf("unexpected filter status %q", filterStatus.Code().String())
		}
		return nil, nil, clustersnapshot.NewFailingPredicateError(pod, filterName, filterReasons, unexpectedErrMsg, p.failingFilterDebugInfo(filterName, nodeInfo))
	}

	return nodeInfo.Node(), state, nil
}

// RunReserveOnNode runs the scheduler framework Reserve phase to update the scheduler plugins state to reflect the Pod being scheduled on the Node.
func (p *SchedulerPluginRunner) RunReserveOnNode(pod *apiv1.Pod, nodeName string, postFilterState *schedulerframework.CycleState) error {
	p.fwHandle.DelegatingLister.UpdateDelegate(p.snapshot)
	defer p.fwHandle.DelegatingLister.ResetDelegate()

	status := p.fwHandle.Framework.RunReservePluginsReserve(context.Background(), postFilterState, pod, nodeName)
	if !status.IsSuccess() {
		return fmt.Errorf("couldn't reserve node %s for pod %s/%s: %v", nodeName, pod.Namespace, pod.Name, status.Message())
	}
	return nil
}

func (p *SchedulerPluginRunner) failingFilterDebugInfo(filterName string, nodeInfo *framework.NodeInfo) string {
	infoParts := []string{fmt.Sprintf("nodeName: %q", nodeInfo.Node().Name)}

	switch filterName {
	case "TaintToleration":
		infoParts = append(infoParts, fmt.Sprintf("nodeTaints: %#v", nodeInfo.Node().Spec.Taints))
	}

	return strings.Join(infoParts, ", ")
}
