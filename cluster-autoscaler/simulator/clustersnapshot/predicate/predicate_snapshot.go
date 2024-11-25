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
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/framework"
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
		result = append(result, framework.WrapSchedulerNodeInfo(schedNodeInfo, nil, nil))
	}
	return result, nil
}

// AddNodeInfo adds the provided internal NodeInfo to the snapshot.
func (s *PredicateSnapshot) AddNodeInfo(nodeInfo *framework.NodeInfo) error {
	return s.SnapshotBase.AddSchedulerNodeInfo(nodeInfo.ToScheduler())
}

// RemoveNodeInfo removes a NodeInfo matching the provided nodeName from the snapshot.
func (s *PredicateSnapshot) RemoveNodeInfo(nodeName string) error {
	return s.SnapshotBase.RemoveSchedulerNodeInfo(nodeName)
}

// SchedulePod adds pod to the snapshot and schedules it to given node.
func (s *PredicateSnapshot) SchedulePod(pod *apiv1.Pod, nodeName string) clustersnapshot.SchedulingError {
	if schedErr := s.pluginRunner.RunFiltersOnNode(pod, nodeName); schedErr != nil {
		return schedErr
	}
	if err := s.ForceAddPod(pod, nodeName); err != nil {
		return clustersnapshot.NewSchedulingInternalError(pod, err.Error())
	}
	return nil
}

// SchedulePodOnAnyNodeMatching adds pod to the snapshot and schedules it to any node matching the provided function.
func (s *PredicateSnapshot) SchedulePodOnAnyNodeMatching(pod *apiv1.Pod, anyNodeMatching func(*framework.NodeInfo) bool) (string, clustersnapshot.SchedulingError) {
	nodeName, schedErr := s.pluginRunner.RunFiltersUntilPassingNode(pod, anyNodeMatching)
	if schedErr != nil {
		return "", schedErr
	}
	if err := s.ForceAddPod(pod, nodeName); err != nil {
		return "", clustersnapshot.NewSchedulingInternalError(pod, err.Error())
	}
	return nodeName, nil
}

// UnschedulePod removes the given Pod from the given Node inside the snapshot.
func (s *PredicateSnapshot) UnschedulePod(namespace string, podName string, nodeName string) error {
	return s.ForceRemovePod(namespace, podName, nodeName)
}

// CheckPredicates checks whether scheduler predicates pass for the given pod on the given node.
func (s *PredicateSnapshot) CheckPredicates(pod *apiv1.Pod, nodeName string) clustersnapshot.SchedulingError {
	return s.pluginRunner.RunFiltersOnNode(pod, nodeName)
}
