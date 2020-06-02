package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// LoadData loads all the system information and returns it in a struct
func LoadData() (BasicInfo, error) {
	var data BasicInfo

	percentage, err := cpu.Percent(1, true)
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

	return data, nil
}
