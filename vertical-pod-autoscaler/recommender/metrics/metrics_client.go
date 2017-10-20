/*
Copyright 2017 The Kubernetes Authors.

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

package metrics

import (
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	v1lister "k8s.io/kubernetes/pkg/client/listers/core/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1alpha1"
	resourceclient "k8s.io/metrics/pkg/client/clientset_generated/clientset/typed/metrics/v1alpha1"
)

//Client providing utilization metrics and metadata for all containers within cluster
type Client interface {
	// Returns ContainerUtilizationSnapshot for each container in the cluster.
	// Snapshot window depends on the configuration of metrics server.
	GetContainersUtilization() ([]*ContainerUtilizationSnapshot, error)
}

type metricsClient struct {
	metricsGetter   resourceclient.PodMetricsesGetter
	podLister       v1lister.PodLister
	namespaceLister v1lister.NamespaceLister
}

// NewClient creates new metrics client, which can be used to calculate containers utilization.
// It requires PodLister, NamespaceLister for get containers metadata and PodMetricsesGetter to get actual metrics.
func NewClient(metricsGetter resourceclient.PodMetricsesGetter, podLister v1lister.PodLister, namespaceLister v1lister.NamespaceLister) Client {
	client := &metricsClient{
		metricsGetter:   metricsGetter,
		podLister:       podLister,
		namespaceLister: namespaceLister,
	}
	glog.V(3).Infof("New metricsClient created %+v", client)

	return client
}

func (client *metricsClient) GetContainersUtilization() ([]*ContainerUtilizationSnapshot, error) {
	glog.V(3).Infof("Getting ContainersUtilization")

	containerSpecs, err := client.getContainersSpec()
	if err != nil {
		return nil, err
	}
	glog.V(3).Infof("%v containerSpecs retrived", len(containerSpecs))

	metricsSnapshots, err := client.getContainersMetrics()
	if err != nil {
		return nil, err
	}
	glog.V(3).Infof("%v metricsSnapshots retrived", len(metricsSnapshots))

	utilizationSnapshots, err := calculateUtilization(metricsSnapshots, containerSpecs)
	if err != nil {
		return nil, err
	}
	glog.V(3).Infof("%v utilizationSnapshots calculated", len(utilizationSnapshots))

	return utilizationSnapshots, nil
}

func (client *metricsClient) getContainersSpec() ([]*basicContainerSpec, error) {
	var containerSpecs []*basicContainerSpec

	pods, err := client.podLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	for _, pod := range pods {
		for _, container := range pod.Spec.Containers {
			containerSpec := newContainerSpec(container, pod)
			containerSpecs = append(containerSpecs, containerSpec)
		}
	}

	return containerSpecs, nil
}

func (client *metricsClient) getContainersMetrics() ([]*containerMetricsSnapshot, error) {
	var metricsSnapshots []*containerMetricsSnapshot

	namespaces, err := client.getAllNamespaces()
	if err != nil {
		return nil, err
	}
	glog.V(3).Infof("%v namespaces retrived: %+v", len(namespaces), namespaces)

	for _, namespace := range namespaces {
		podMetricsInterface := client.metricsGetter.PodMetricses(namespace)
		podMetricsList, err := podMetricsInterface.List(metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		glog.V(3).Infof("podMetricsList retrived for: %+v", namespace)
		for _, podMetrics := range podMetricsList.Items {
			containerSnapshots := createContainerMetricsSnapshots(podMetrics)
			metricsSnapshots = append(metricsSnapshots, containerSnapshots...)
		}
	}

	return metricsSnapshots, nil
}

func createContainerMetricsSnapshots(podMetrics v1alpha1.PodMetrics) []*containerMetricsSnapshot {
	snapshots := make([]*containerMetricsSnapshot, len(podMetrics.Containers))
	for i, containerMetrics := range podMetrics.Containers {
		snapshots[i] = newContainerMetricsSnapshot(containerMetrics, podMetrics)
	}
	return snapshots
}

func calculateUtilization(snapshots []*containerMetricsSnapshot, specifications []*basicContainerSpec) ([]*ContainerUtilizationSnapshot, error) {
	specsMap := make(map[containerID]*basicContainerSpec, len(specifications))
	for _, spec := range specifications {
		specsMap[spec.ID] = spec
	}

	result := make([]*ContainerUtilizationSnapshot, len(snapshots))

	for i, snap := range snapshots {
		spec := specsMap[snap.ID]
		utilizationSnapshot, err := NewContainerUtilizationSnapshot(snap, spec)
		if err != nil {
			return nil, err
		}
		result[i] = utilizationSnapshot
	}

	return result, nil
}

func newContainerSpec(container v1.Container, pod *v1.Pod) *basicContainerSpec {
	return &basicContainerSpec{
		ID: containerID{
			PodName:       pod.Name,
			Namespace:     pod.Namespace,
			ContainerName: container.Name,
		},
		PodLabels:    pod.Labels,
		CreationTime: pod.CreationTimestamp,
		Image:        container.Image,
		Request:      container.Resources.Requests,
	}
}

func newContainerMetricsSnapshot(containerMetrics v1alpha1.ContainerMetrics, podMetrics v1alpha1.PodMetrics) *containerMetricsSnapshot {
	return &containerMetricsSnapshot{
		ID: containerID{
			ContainerName: containerMetrics.Name,
			Namespace:     podMetrics.Namespace,
			PodName:       podMetrics.Name,
		},
		Usage:          containerMetrics.Usage,
		SnapshotTime:   podMetrics.Timestamp,
		SnapshotWindow: podMetrics.Window,
	}
}

func (client *metricsClient) getAllNamespaces() ([]string, error) {
	namespaces, err := client.namespaceLister.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	count := len(namespaces)
	result := make([]string, count)

	for i, namespace := range namespaces {
		result[i] = namespace.Name
	}

	return result, nil
}
