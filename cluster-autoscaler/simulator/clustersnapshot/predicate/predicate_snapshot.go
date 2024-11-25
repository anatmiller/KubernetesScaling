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
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	drautils "k8s.io/autoscaler/cluster-autoscaler/simulator/dynamicresources/utils"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/framework"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

// PredicateSnapshot implements ClusterSnapshot on top of a SnapshotBase by using
// SchedulerBasedPredicateChecker to check scheduler predicates.
type PredicateSnapshot struct {
	clustersnapshot.SnapshotBase
	pluginRunner *SchedulerPluginRunner
	draEnabled   bool
}

// NewPredicateSnapshot builds a PredicateSnapshot.
func NewPredicateSnapshot(snapshotBase clustersnapshot.SnapshotBase, fwHandle *framework.Handle, draEnabled bool) *PredicateSnapshot {
	snapshot := &PredicateSnapshot{
		SnapshotBase: snapshotBase,
		draEnabled:   draEnabled,
	}
	snapshot.pluginRunner = NewSchedulerPluginRunner(fwHandle, snapshot)
	return snapshot
}

// GetNodeInfo returns an internal NodeInfo wrapping the relevant schedulerframework.NodeInfo.
func (s *PredicateSnapshot) GetNodeInfo(nodeName string) (*framework.NodeInfo, error) {
	schedNodeInfo, err := s.SnapshotBase.NodeInfos().Get(nodeName)
	if err != nil {
		return nil, err
	}
	if s.draEnabled {
		return s.SnapshotBase.DraSnapshot().WrapSchedulerNodeInfo(schedNodeInfo)
	}
	return framework.WrapSchedulerNodeInfo(schedNodeInfo, nil, nil), nil
}

// ListNodeInfos returns internal NodeInfos wrapping all schedulerframework.NodeInfos in the snapshot.
func (s *PredicateSnapshot) ListNodeInfos() ([]*framework.NodeInfo, error) {
	schedNodeInfos, err := s.SnapshotBase.NodeInfos().List()
	if err != nil {
		return nil, err
	}
	var result []*framework.NodeInfo
	for _, schedNodeInfo := range schedNodeInfos {
		if s.draEnabled {
			nodeInfo, err := s.SnapshotBase.DraSnapshot().WrapSchedulerNodeInfo(schedNodeInfo)
			if err != nil {
				return nil, err
			}
			result = append(result, nodeInfo)
		} else {
			result = append(result, framework.WrapSchedulerNodeInfo(schedNodeInfo, nil, nil))
		}
	}
	return result, nil
}

// AddNodeInfo adds the provided internal NodeInfo to the snapshot.
func (s *PredicateSnapshot) AddNodeInfo(nodeInfo *framework.NodeInfo) error {
	if s.draEnabled && len(nodeInfo.LocalResourceSlices) > 0 {
		err := s.SnapshotBase.DraSnapshot().AddNodeResourceSlices(nodeInfo.Node().Name, nodeInfo.LocalResourceSlices)
		if err != nil {
			return fmt.Errorf("couldn't add ResourceSlices to DRA snapshot: %v", err)
		}

		for _, podInfo := range nodeInfo.Pods() {
			err := s.SnapshotBase.DraSnapshot().AddClaims(podInfo.NeededResourceClaims)
			if err != nil {
				return fmt.Errorf("couldn't add ResourceSlices to DRA snapshot: %v", err)
			}
		}
	}

	return s.SnapshotBase.AddSchedulerNodeInfo(nodeInfo.ToScheduler())
}

// RemoveNodeInfo removes a NodeInfo matching the provided nodeName from the snapshot.
func (s *PredicateSnapshot) RemoveNodeInfo(nodeName string) error {
	if s.draEnabled {
		nodeInfo, err := s.GetNodeInfo(nodeName)
		if err != nil {
			return err
		}

		s.SnapshotBase.DraSnapshot().RemoveNodeResourceSlices(nodeName)

		for _, pod := range nodeInfo.Pods() {
			s.SnapshotBase.DraSnapshot().RemovePodClaims(pod.Pod)
		}
	}

	return s.SnapshotBase.RemoveSchedulerNodeInfo(nodeName)
}

// SchedulePod adds pod to the snapshot and schedules it to given node.
func (s *PredicateSnapshot) SchedulePod(pod *apiv1.Pod, nodeName string) clustersnapshot.SchedulingError {
	node, cycleState, schedErr := s.pluginRunner.RunFiltersOnNode(pod, nodeName)
	if schedErr != nil {
		return schedErr
	}

	if s.draEnabled {
		if err := s.handleResourceClaimModifications(pod, node, cycleState); err != nil {
			return clustersnapshot.NewSchedulingInternalError(pod, err.Error())
		}
	}

	if err := s.ForceAddPod(pod, nodeName); err != nil {
		return clustersnapshot.NewSchedulingInternalError(pod, err.Error())
	}
	return nil
}

// SchedulePodOnAnyNodeMatching adds pod to the snapshot and schedules it to any node matching the provided function.
func (s *PredicateSnapshot) SchedulePodOnAnyNodeMatching(pod *apiv1.Pod, anyNodeMatching func(*framework.NodeInfo) bool) (string, clustersnapshot.SchedulingError) {
	node, cycleState, schedErr := s.pluginRunner.RunFiltersUntilPassingNode(pod, anyNodeMatching)
	if schedErr != nil {
		return "", schedErr
	}

	if s.draEnabled {
		if err := s.handleResourceClaimModifications(pod, node, cycleState); err != nil {
			return "", clustersnapshot.NewSchedulingInternalError(pod, err.Error())
		}
	}

	if err := s.ForceAddPod(pod, node.Name); err != nil {
		return "", clustersnapshot.NewSchedulingInternalError(pod, err.Error())
	}
	return node.Name, nil
}

// UnschedulePod removes the given Pod from the given Node inside the snapshot.
func (s *PredicateSnapshot) UnschedulePod(namespace string, podName string, nodeName string) error {
	if s.draEnabled {
		nodeInfo, err := s.GetNodeInfo(nodeName)
		if err != nil {
			return err
		}

		var foundPod *apiv1.Pod
		for _, pod := range nodeInfo.Pods() {
			if pod.Namespace == namespace && pod.Name == podName {
				foundPod = pod.Pod
				break
			}
		}
		if foundPod == nil {
			return fmt.Errorf("pod %s/%s not found on node %s", namespace, podName, nodeName)
		}

		if err := s.SnapshotBase.DraSnapshot().UnreservePodClaims(foundPod); err != nil {
			return err
		}
	}

	return s.ForceRemovePod(namespace, podName, nodeName)
}

// CheckPredicates checks whether scheduler predicates pass for the given pod on the given node.
func (s *PredicateSnapshot) CheckPredicates(pod *apiv1.Pod, nodeName string) clustersnapshot.SchedulingError {
	_, _, err := s.pluginRunner.RunFiltersOnNode(pod, nodeName)
	return err
}

func (s *PredicateSnapshot) handleResourceClaimModifications(pod *apiv1.Pod, node *apiv1.Node, postFilterState *schedulerframework.CycleState) error {
	if len(pod.Spec.ResourceClaims) == 0 {
		return nil
	}
	// We need to run the scheduler Reserve phase to allocate the appropriate ResourceClaims in the DRA snapshot. The allocations are
	// actually computed and cached in the Filter phase, and Reserve only grabs them from the cycle state. So this should be quick, but
	// it needs the cycle state from after running the Filter phase.
	err := s.pluginRunner.RunReserveOnNode(pod, node.Name, postFilterState)
	if err != nil {
		return fmt.Errorf("error while trying to run Reserve node %s for pod %s/%s: %v", node.Name, pod.Namespace, pod.Name, err)
	}

	// The pod isn't added to the ReservedFor field of the claim during the Reserve phase (it happens later, in PreBind). We can just do it
	// manually here. It shouldn't fail, it only fails if ReservedFor is at max length already, but that is checked during the Filter phase.
	err = s.SnapshotBase.DraSnapshot().ReservePodClaims(pod)
	if err != nil {
		return fmt.Errorf("couldnn't add pod reservations to claims, this shouldn't happen: %v", err)
	}

	// Verify that all needed claims are tracked in the DRA snapshot, allocated, and available on the Node.
	claims, err := s.SnapshotBase.DraSnapshot().PodClaims(pod)
	if err != nil {
		return fmt.Errorf("couldn't obtain pod %s/%s claims: %v", pod.Namespace, pod.Name, err)
	}
	for _, claim := range claims {
		if available, err := drautils.ClaimAvailableOnNode(claim, node); err != nil || !available {
			return fmt.Errorf("pod %s/%s needs claim %s to schedule, but it isn't available on node %s (allocated: %v, available: %v, err: %v)", pod.Namespace, pod.Name, claim.Name, node.Name, drautils.ClaimAllocated(claim), available, err)
		}
	}
	return nil
}
