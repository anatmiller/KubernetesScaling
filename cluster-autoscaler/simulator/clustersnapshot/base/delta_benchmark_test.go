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

package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/framework"
)

func BenchmarkBuildNodeInfoList(b *testing.B) {
	testCases := []struct {
		nodeCount int
	}{
		{
			nodeCount: 1000,
		},
		{
			nodeCount: 5000,
		},
		{
			nodeCount: 15000,
		},
		{
			nodeCount: 100000,
		},
	}

	for _, tc := range testCases {
		b.Run(fmt.Sprintf("fork add 1000 to %d", tc.nodeCount), func(b *testing.B) {
			nodes := clustersnapshot.CreateTestNodes(tc.nodeCount + 1000)
			clusterSnapshot := NewDeltaSnapshotBase()
			if err := clusterSnapshot.SetClusterState(nodes[:tc.nodeCount], nil); err != nil {
				assert.NoError(b, err)
			}
			clusterSnapshot.Fork()
			for _, node := range nodes[tc.nodeCount:] {
				if err := clusterSnapshot.AddNodeInfo(framework.NewTestNodeInfo(node)); err != nil {
					assert.NoError(b, err)
				}
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				list := clusterSnapshot.data.buildNodeInfoList()
				if len(list) != tc.nodeCount+1000 {
					assert.Equal(b, len(list), tc.nodeCount+1000)
				}
			}
		})
	}
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("base %d", tc.nodeCount), func(b *testing.B) {
			nodes := clustersnapshot.CreateTestNodes(tc.nodeCount)
			clusterSnapshot := NewDeltaSnapshotBase()
			if err := clusterSnapshot.SetClusterState(nodes, nil); err != nil {
				assert.NoError(b, err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				list := clusterSnapshot.data.buildNodeInfoList()
				if len(list) != tc.nodeCount {
					assert.Equal(b, len(list), tc.nodeCount)
				}
			}
		})
	}
}
