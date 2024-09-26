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

package provider

import (
	resourceapi "k8s.io/api/resource/v1beta1"
	drasnapshot "k8s.io/autoscaler/cluster-autoscaler/simulator/dynamicresources/snapshot"
	"k8s.io/client-go/informers"
)

// Provider provides DRA-related objects.
type Provider struct {
	resourceClaims providerClaimLister
	resourceSlices providerSliceLister
	deviceClasses  providerClassLister
}

// NewProviderFromInformers returns a new Provider which uses InformerFactory listers to list the DRA resources.
func NewProviderFromInformers(informerFactory informers.SharedInformerFactory) *Provider {
	claims := &resourceClaimApiLister{apiLister: informerFactory.Resource().V1beta1().ResourceClaims().Lister()}
	slices := &resourceSliceApiLister{apiLister: informerFactory.Resource().V1beta1().ResourceSlices().Lister()}
	devices := &deviceClassApiLister{apiLister: informerFactory.Resource().V1beta1().DeviceClasses().Lister()}
	return NewProvider(claims, slices, devices)
}

// NewProvider returns a new Provider which uses the provided listers to list the DRA resources.
func NewProvider(claims providerClaimLister, slices providerSliceLister, classes providerClassLister) *Provider {
	return &Provider{
		resourceClaims: claims,
		resourceSlices: slices,
		deviceClasses:  classes,
	}
}

// Snapshot returns a snapshot of all DRA resources at a ~single point in time.
func (p *Provider) Snapshot() (drasnapshot.Snapshot, error) {
	claims, err := p.resourceClaims.List()
	if err != nil {
		return drasnapshot.Snapshot{}, err
	}
	claimMap := make(map[drasnapshot.ResourceClaimId]*resourceapi.ResourceClaim)
	for _, claim := range claims {
		claimMap[drasnapshot.ResourceClaimId{Name: claim.Name, Namespace: claim.Namespace}] = claim
	}

	slices, err := p.resourceSlices.List()

	if err != nil {
		return drasnapshot.Snapshot{}, err
	}
	slicesMap := make(map[string][]*resourceapi.ResourceSlice)
	var nonNodeLocalSlices []*resourceapi.ResourceSlice
	for _, slice := range slices {
		if slice.Spec.NodeName == "" {
			nonNodeLocalSlices = append(nonNodeLocalSlices, slice)
		} else {
			slicesMap[slice.Spec.NodeName] = append(slicesMap[slice.Spec.NodeName], slice)
		}
	}

	classes, err := p.deviceClasses.List()
	if err != nil {
		return drasnapshot.Snapshot{}, err
	}
	classMap := make(map[string]*resourceapi.DeviceClass)
	for _, class := range classes {
		classMap[class.Name] = class
	}

	return drasnapshot.NewSnapshot(claimMap, slicesMap, nonNodeLocalSlices, classMap), nil
}
