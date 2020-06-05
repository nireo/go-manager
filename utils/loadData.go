package utils

import (
	"fmt"
	"log"
	"strings"
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

	tempProcesses := LoadAllProcesses()
	var processes [][]string

	// add the header row to the list
	processes = append(processes, []string{"PID", "NAME", "USER", "CPU%", "RUNNING"})

	for index := range tempProcesses {
		singleProcess := ChangeProcessToTableFormat(tempProcesses[index])
		processes = append(processes, singleProcess)
	}

	data.Processes = processes

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
	cpuPercentage, err := process.CPUPercent()
	if err != nil {
		log.Fatalf("failed to get cpu percentage: %v", err)
	}

	isRunning, err := process.IsRunning()
	if err != nil {
		log.Fatalf("failed to get running status: %v", err)
	}

	user, err := process.Username()
	if err != nil {
		log.Fatalf("failed to get username: %v", err)
	}

	name, err := process.Name()
	if err != nil {
		log.Fatalf("failed to get name: %v", err)
	}

	newProcessEntry := Process{
		Pid:           process.Pid,
		CPUPercentage: cpuPercentage,
		Running:       isRunning,
		User:          user,
		Name:          name,
	}

	return newProcessEntry
}

// ChangeProcessToTableFormat returns all the process fields in an array of string
func ChangeProcessToTableFormat(p Process) []string {
	isRunningString := fmt.Sprintf("%t", p.Running)
	isRunningString = strings.ToUpper(isRunningString)

	var tableFormat []string
	tableFormat = append(tableFormat, fmt.Sprintf("%d", p.Pid))
	tableFormat = append(tableFormat, p.Name)
	tableFormat = append(tableFormat, p.User)
	tableFormat = append(tableFormat, fmt.Sprintf("%.2f", p.CPUPercentage))
	tableFormat = append(tableFormat, fmt.Sprintf("[%s](fg:green)", isRunningString))

	return tableFormat
}
