package podsharding

import (
	"fmt"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/autoscaler/cluster-autoscaler/context"
)

// PredicatePodShardFilter implements PodShardFilter interface. Initial set of pods is based on list of pod UIDs stored in
// PodShard. Then list is extended by testing scheduler predicates for other pods on test NodeInfos build via cloudprovider
// based on PodShard.NodeGroupDescriptor.
type PredicatePodShardFilter struct{}

// NewPredicatePodShardFilter creates an instance of PredicatePodShardFilter
func NewPredicatePodShardFilter() *PredicatePodShardFilter {
	return &PredicatePodShardFilter{}
}

// FilterPods filters pod list against PodShard
func (p *PredicatePodShardFilter) FilterPods(context *context.AutoscalingContext, selectedPodShard *PodShard, allPodShards []*PodShard, pods []*apiv1.Pod) (PodFilteringResult, error) {
	podsByUid := make(map[types.UID]*apiv1.Pod)
	for _, pod := range pods {
		podsByUid[pod.UID] = pod
	}

	if len(selectedPodShard.PodUids) < 1 {
		return PodFilteringResult{}, fmt.Errorf("not enough pods associated to the selected PodShard")
	}

	// list of shards for which we want to have
	finalPodShards := make(map[ShardSignature]bool)

	// add selected shard to list of final shards
	finalPodShards[selectedPodShard.Signature()] = true

	// iterate over all selected shards and build final set of Pod UIDs
	selectedPodUids := make(map[types.UID]bool)
	for _, shard := range allPodShards {
		if finalPodShards[shard.Signature()] {
			for podUid := range shard.PodUids {
				selectedPodUids[podUid] = true
			}
		}
	}

	// translate UIDs of selected Pods to Pods
	var selectedPods []*apiv1.Pod
	for _, pod := range pods {
		if selectedPodUids[pod.UID] {
			selectedPods = append(selectedPods, pod)
		}
	}

	return PodFilteringResult{
		Pods: selectedPods,
	}, nil
}
