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

package snapshot

import (
	"fmt"

	resourceapi "k8s.io/api/resource/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	drautils "k8s.io/autoscaler/cluster-autoscaler/simulator/dynamicresources/utils"
	"k8s.io/dynamic-resource-allocation/structured"
)

type snapshotClaimTracker Snapshot

func (s snapshotClaimTracker) List() ([]*resourceapi.ResourceClaim, error) {
	var result []*resourceapi.ResourceClaim
	for _, claim := range s.resourceClaimsById {
		result = append(result, claim)
	}
	return result, nil
}

func (s snapshotClaimTracker) Get(namespace, claimName string) (*resourceapi.ResourceClaim, error) {
	claim, found := s.resourceClaimsById[ResourceClaimId{Name: claimName, Namespace: namespace}]
	if !found {
		return nil, fmt.Errorf("claim %s/%s not found", namespace, claimName)
	}
	return claim, nil
}

func (s snapshotClaimTracker) ListAllAllocatedDevices() (sets.Set[structured.DeviceID], error) {
	result := sets.New[structured.DeviceID]()
	for _, claim := range s.resourceClaimsById {
		result = result.Union(drautils.ClaimAllocatedDevices(claim))
	}
	return result, nil
}

func (s snapshotClaimTracker) SignalClaimPendingAllocation(claimUid types.UID, allocatedClaim *resourceapi.ResourceClaim) error {
	ref := ResourceClaimId{Name: allocatedClaim.Name, Namespace: allocatedClaim.Namespace}
	claim, found := s.resourceClaimsById[ref]
	if !found {
		return fmt.Errorf("claim %s/%s not found", allocatedClaim.Namespace, allocatedClaim.Name)
	}
	if claim.UID != claimUid {
		return fmt.Errorf("claim %s/%s: snapshot has UID %q, allocation came for UID %q - shouldn't happenn", allocatedClaim.Namespace, allocatedClaim.Name, claim.UID, claimUid)
	}
	s.resourceClaimsById[ref] = allocatedClaim
	return nil
}

func (s snapshotClaimTracker) ClaimHasPendingAllocation(claimUid types.UID) bool {
	return false
}

func (s snapshotClaimTracker) RemoveClaimPendingAllocation(claimUid types.UID) (deleted bool) {
	// This method is only called during the Bind phase of scheduler framework, which is never run by CA. We need to implement
	// it to satisfy the interface, but it should never be called.
	panic("snapshotClaimTracker.RemoveClaimPendingAllocation() was called - this should never happen")
}

func (s snapshotClaimTracker) AssumeClaimAfterAPICall(claim *resourceapi.ResourceClaim) error {
	// This method is only called during the Bind phase of scheduler framework, which is never run by CA. We need to implement
	// it to satisfy the interface, but it should never be called.
	panic("snapshotClaimTracker.AssumeClaimAfterAPICall() was called - this should never happen")
}

func (s snapshotClaimTracker) AssumedClaimRestore(namespace, claimName string) {
	// This method is only called during the Bind phase of scheduler framework, which is never run by CA. We need to implement
	// it to satisfy the interface, but it should never be called.
	panic("snapshotClaimTracker.AssumedClaimRestore() was called - this should never happen")
}
