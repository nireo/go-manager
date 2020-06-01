package main

import (
	"fmt"
	"time"

	"github.com/nireo/go-manager/ui"
	"github.com/nireo/go-manager/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var data utils.BasicInfo

func main() {
	fmt.Printf("%d", utils.GetTerminalWidth())

	duration, err := time.ParseDuration("2s")
	if err != nil {
		panic(err)
	}

	for {
		percentage, err := cpu.Percent(1, true)
		if err != nil {
			panic(err)
		}

		memory, err := mem.SwapMemory()
		if err != nil {
			panic(err)
		}

		data.CorePercentages = percentage
		data.MemoryUsed = memory.Used
		data.TotalMemory = memory.Total
		data.MemoryUsedPercent = memory.UsedPercent

		ui.InitUI(data)
		time.Sleep(duration)
	}
}
