package utils

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	resourceapi "k8s.io/api/resource/v1beta1"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/framework"
)

// CalculateDynamicResourceUtils calculates a map of ResourceSlice pool utilization grouped by the driver and pool. Returns
// an error if the NodeInfo doesn't have all ResourceSlices from a pool.
func CalculateDynamicResourceUtils(nodeInfo *framework.NodeInfo) (map[string]map[string]float64, error) {
	result := map[string]map[string]float64{}
	claims := NodeInfoResourceClaims(nodeInfo)
	allocatedDevices, err := GroupAllocatedDevices(claims)
	if err != nil {
		return nil, err
	}
	for driverName, slicesByPool := range GroupSlices(nodeInfo.LocalResourceSlices) {
		result[driverName] = map[string]float64{}
		for poolName, poolSlices := range slicesByPool {
			currentSlices, err := AllCurrentGenSlices(poolSlices)
			if err != nil {
				return nil, fmt.Errorf("pool %q error: %v", poolName, err)
			}
			poolDevices := GetAllDevices(currentSlices)
			allocatedDeviceNames := allocatedDevices[driverName][poolName]
			unallocated, allocated := splitDevicesByAllocation(poolDevices, allocatedDeviceNames)
			result[driverName][poolName] = calculatePoolUtil(unallocated, allocated)
		}
	}
	return result, nil
}

// HighestDynamicResourceUtil returns the ResourceSlice driver and pool with the highest utilization.
func HighestDynamicResourceUtil(nodeInfo *framework.NodeInfo) (v1.ResourceName, float64, error) {
	utils, err := CalculateDynamicResourceUtils(nodeInfo)
	if err != nil {
		return "", 0, err
	}

	highestUtil := 0.0
	var highestResourceName v1.ResourceName
	for driverName, utilsByPool := range utils {
		for poolName, util := range utilsByPool {
			if util >= highestUtil {
				highestUtil = util
				highestResourceName = v1.ResourceName(driverName + "/" + poolName)
			}
		}
	}
	return highestResourceName, highestUtil, nil
}

func calculatePoolUtil(unallocated, allocated []resourceapi.Device) float64 {
	numAllocated := float64(len(allocated))
	numUnallocated := float64(len(unallocated))
	return numAllocated / (numAllocated + numUnallocated)
}

func splitDevicesByAllocation(devices []resourceapi.Device, allocatedNames []string) (unallocated, allocated []resourceapi.Device) {
	allocatedNamesSet := map[string]bool{}
	for _, allocatedName := range allocatedNames {
		allocatedNamesSet[allocatedName] = true
	}
	for _, device := range devices {
		if allocatedNamesSet[device.Name] {
			allocated = append(allocated, device)
		} else {
			unallocated = append(unallocated, device)
		}
	}
	return unallocated, allocated
}
