package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	envContainersPerGPU      = "DP_CONTAINERS_PER_GPU"
	defaultContainersPerGPU  = 1
	virtualDeviceIDSeperator = "-VGPU-"
)

func getNumOfContainersPerGPU() uint {
	numStr, present := os.LookupEnv(envContainersPerGPU)

	if !present {
		return defaultContainersPerGPU
	}

	rawContainersPerGPU, err := strconv.Atoi(numStr)
	if err != nil {
		log.Panicf("parse env DP_CONTAINERS_PER_GPU error")
	}
	if rawContainersPerGPU < 1 {
		log.Panicf("DP_CONTAINERS_PER_GPU must be greater than 1")
	}

	containersPerGPU := uint(rawContainersPerGPU)
	return containersPerGPU
}

func generateVirtualGPUDeviceID(deviceID string, virtualID uint) string {
	return fmt.Sprintf("%s%s%d", deviceID, virtualDeviceIDSeperator, virtualID)
}

func getRealDeviceID(virtualDeviceID string) string {
	return strings.Split(virtualDeviceID, virtualDeviceIDSeperator)[0]
}
