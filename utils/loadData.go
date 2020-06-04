package utils

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

// LoadData loads all the system information and returns it in a struct
func LoadData() (BasicInfo, error) {
	var data BasicInfo

	duration, err := time.ParseDuration("1s")
	percentage, err := cpu.Percent(duration, true)
	if err != nil {
		return data, err
	}

	memory, err := mem.SwapMemory()
	if err != nil {
		return data, err
	}

	data.CorePercentages = percentage
	data.MemoryUsed = memory.Used
	data.TotalMemory = memory.Total
	data.MemoryUsedPercent = memory.UsedPercent

	info, err := host.Info()
	if err != nil {
		return data, err
	}

	data.Hostname = info.Hostname
	data.Uptime = info.Uptime
	data.KernelVersion = info.KernelVersion
	data.Procs = info.Procs

	return data, nil
}

// LoadAllProcesses returns a filtered list of processes which are
func LoadAllProcesses() []Process {
	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}

	var filteredProcesses []Process
	for index := range processes {
		filteredProcesses = append(filteredProcesses, LoadSingleProcessData(processes[index]))
	}

	return filteredProcesses
}

// LoadSingleProcessData loads all the data into the predefined process struct
func LoadSingleProcessData(process *process.Process) Process {
	exe, err := process.Exe()
	if err != nil {
		panic(err)
	}

	cpuPercentage, err := process.CPUPercent()
	if err != nil {
		panic(err)
	}

	isRunning, err := process.IsRunning()
	if err != nil {
		panic(err)
	}

	user, err := process.Username()
	if err != nil {
		panic(err)
	}

	name, err := process.Name()
	if err != nil {
		panic(err)
	}

	newProcessEntry := Process{
		Pid:           process.Pid,
		Exe:           exe,
		CPUPercentage: cpuPercentage,
		Running:       isRunning,
		User:          user,
		Name:          name,
	}

	return newProcessEntry
}
