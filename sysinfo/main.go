package sysinfo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// GetHostInfo retrieves basic host information.
func GetHostInfo() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error retrieving hostname: %v", err)
	}

	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Get kernel version using `uname -r`
	kernelVersion, err := exec.Command("uname", "-r").Output()
	if err != nil {
		log.Fatalf("Error retrieving kernel version: %v", err)
	}

	// Get OS version using `sw_vers -productVersion` (macOS specific)
	osVersion := "Unknown"
	if osName == "darwin" { // macOS
		versionOutput, err := exec.Command("sw_vers", "-productVersion").Output()
		if err != nil {
			log.Printf("Error retrieving OS version: %v", err)
		} else {
			osVersion = strings.TrimSpace(string(versionOutput))
		}
	}

	return fmt.Sprintf(
		"Hostname: %s\nOS: %s\nOS Version: %s\nKernel Version: %s\nArchitecture: %s\n",
		hostname, osName, osVersion, strings.TrimSpace(string(kernelVersion)), arch,
	)
}

// GetCPUInfo retrieves CPU usage information.
func GetCPUInfo() string {
	// Use `top` or `ps` commands to retrieve CPU usage
	cpuUsage, err := exec.Command("sh", "-c", "top -l 1 | grep 'CPU usage'").Output()
	if err != nil {
		log.Fatalf("Error retrieving CPU info: %v", err)
	}

	return fmt.Sprintf("CPU Info: %s", strings.TrimSpace(string(cpuUsage)))
}

// GetMemoryInfo retrieves memory usage information.
func GetMemoryInfo() string {
	// Use `vm_stat` to retrieve memory information on macOS
	memStats, err := exec.Command("vm_stat").Output()
	if err != nil {
		log.Fatalf("Error retrieving memory info: %v", err)
	}

	// Parse `vm_stat` output to calculate memory usage
	lines := strings.Split(string(memStats), "\n")
	pageSize := 4096 // macOS default page size in bytes
	var totalPages, freePages int

	for _, line := range lines {
		if strings.Contains(line, "Pages free") {
			fmt.Sscanf(line, "Pages free: %d.", &freePages)
		}
		if strings.Contains(line, "Pages active") || strings.Contains(line, "Pages inactive") || strings.Contains(line, "Pages speculative") || strings.Contains(line, "Pages wired down") {
			var pages int
			fmt.Sscanf(line, "%*s %*s %d.", &pages)
			totalPages += pages
		}
	}

	totalMemory := totalPages * pageSize
	freeMemory := freePages * pageSize
	usedMemory := totalMemory - freeMemory

	return fmt.Sprintf(
		"Total Memory: %d bytes\nFree Memory: %d bytes\nUsed Memory: %d bytes\n",
		totalMemory, freeMemory, usedMemory,
	)
}
