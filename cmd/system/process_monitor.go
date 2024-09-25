package process_monitor

import (
	find_process "saruman/cmd/system/find"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type SystemStats struct {
	ProcessCpu            float64
	ProcessMemory         float64
	TotalSystemMemory     float64
	AvailableSystemMemory float64
	NumCpu                int
}

const (
	BytesInGB = 1024 * 1024 * 1024
)

func CollectSystemStats(processName string) (*SystemStats, error) {
	pid, err := find_process.ByName(processName)

	result, err := getSystemStats(pid)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func getSystemStats(pid int32) (*SystemStats, error) {
	proc, err := systemProcess(pid)
	if err != nil {
		return nil, err
	}

	processCpu, _ := processCpuUsage(proc)
	processMemory, _ := processMemoryUsage(proc)

	totalSystemMemory, availableSystemMemory, numCpu, _ := systemVirutalMemoryUsage()

	stats := &SystemStats{
		ProcessCpu:            processCpu,
		ProcessMemory:         processMemory,
		TotalSystemMemory:     totalSystemMemory,
		AvailableSystemMemory: availableSystemMemory,
		NumCpu:                numCpu,
	}

	return stats, nil
}

func processCpuUsage(proc *process.Process) (float64, error) {
	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		return 0.0, err
	}

	return cpuPercent, nil
}

func processMemoryUsage(proc *process.Process) (float64, error) {
	memInfo, err := proc.MemoryInfo()
	if err != nil {
		return 0.0, err
	}

	return float64(memInfo.RSS) / (BytesInGB), nil
}

func systemVirutalMemoryUsage() (float64, float64, int, error) {
	vMem, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, 0.0, 0, err
	}

	totalGiB := float64(vMem.Total) / (BytesInGB)
	availableGiB := float64(vMem.Available) / (BytesInGB)

	numCPUs, err := cpu.Counts(true)
	if err != nil {
		return 0.0, 0.0, 0, err
	}

	return totalGiB, availableGiB, numCPUs, nil
}

func systemProcess(pid int32) (*process.Process, error) {
	proc, err := process.NewProcess(pid)

	if err != nil {
		return nil, err
	}

	return proc, nil
}
