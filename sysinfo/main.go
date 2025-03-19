package sysinfo

import (
	"fmt"
	"log"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
)

// GetHostInfo retrieves basic host information.
func GetHostInfo() string {
	host, err := sysinfo.Host()
	if err != nil {
		log.Fatalf("Error retrieving host info: %v", err)
	}

	info := host.Info()
	return fmt.Sprintf(
		"Hostname: %s\nOS: %s %s (%s)\nKernel Version: %s\nArchitecture: %s\n",
		info.Hostname, info.OS.Name, info.OS.Version, info.OS.Platform, info.KernelVersion, info.Architecture,
	)
}

// GetCPUInfo retrieves CPU usage information.
func GetCPUInfo() string {
	host, err := sysinfo.Host()
	if err != nil {
		log.Fatalf("Error retrieving host info: %v", err)
	}

	if cpuTimer, ok := host.(types.CPUTimer); ok {
		cpuInfo, err := cpuTimer.CPUTime()
		if err != nil {
			log.Fatalf("Error retrieving CPU info: %v", err)
		}
		return fmt.Sprintf(
			"CPU User Time: %f seconds\nCPU System Time: %f seconds\nCPU Idle Time: %f seconds\n",
			cpuInfo.User.Seconds(), cpuInfo.System.Seconds(), cpuInfo.Idle.Seconds(),
		)
	}
	return "CPU information not available on this platform.\n"
}

// GetMemoryInfo retrieves memory usage information.
func GetMemoryInfo() string {
	host, err := sysinfo.Host()
	if err != nil {
		log.Fatalf("Error retrieving host info: %v", err)
	}

	memInfo, err := host.Memory()
	if err != nil {
		log.Fatalf("Error retrieving memory info: %v", err)
	}
	return fmt.Sprintf(
		"Total Memory: %d bytes\nFree Memory: %d bytes\nUsed Memory: %d bytes\n",
		memInfo.Total, memInfo.Free, memInfo.Total-memInfo.Free,
	)
}
