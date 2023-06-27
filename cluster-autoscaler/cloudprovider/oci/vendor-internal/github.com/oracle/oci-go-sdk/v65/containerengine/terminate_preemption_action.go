// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TerminatePreemptionAction Terminates the preemptible instance when it is interrupted for eviction.
type TerminatePreemptionAction struct {

	// Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified.
	IsPreserveBootVolume *bool `mandatory:"false" json:"isPreserveBootVolume"`
}

func (m TerminatePreemptionAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TerminatePreemptionAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TerminatePreemptionAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTerminatePreemptionAction TerminatePreemptionAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTerminatePreemptionAction
	}{
		"TERMINATE",
		(MarshalTypeTerminatePreemptionAction)(m),
	}

	return json.Marshal(&s)
}
