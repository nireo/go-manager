package utils

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// BasicInfo has all the the values in the upper box
type BasicInfo struct {
	CorePercentages   []float64
	MemoryUsed        uint64
	TotalMemory       uint64
	MemoryUsedPercent float64
	Hostname          string
	Uptime            uint64
	KernelVersion     string
	Procs             uint64
	Processes         [][]string
}

// Process includes all the types for the table view
type Process struct {
	Pid           int32
	Name          string
	CPUPercentage float64
	Running       bool
	User          string
}

// View data structure
type View struct {
	List            *widgets.List
	SystemInfoList  *widgets.List
	ProcessesWindow *widgets.Table
	Data            BasicInfo
}

// Init initializes the ui
func (view *View) Init() {
	view.List.Title = "CPU & Memory information"
	view.List.SetRect(0, 0, int(GetTerminalWidth()/2), int(GetTerminalHeight()/4))

	view.SystemInfoList.Title = "System information"
	view.SystemInfoList.SetRect(int(GetTerminalWidth()/2), 0, int(GetTerminalWidth()), int(GetTerminalHeight()/4))

	view.ProcessesWindow.Title = "Processes"
	view.ProcessesWindow.SetRect(0, 13, int(GetTerminalWidth()), int(GetTerminalHeight()/3)*2)
	view.ProcessesWindow.Rows = [][]string{
		[]string{"PID", "NAME", "USER", "CPU%", "RUNNING"},
	}
}

// NewView returns a pointer to an view struct
func NewView() *View {
	view := &View{
		List:            widgets.NewList(),
		SystemInfoList:  widgets.NewList(),
		ProcessesWindow: widgets.NewTable(),
		Data:            BasicInfo{},
	}

	view.Init()
	return view
}

// Resize is used when an terminal resize event occurs and it updates list dimensions accordinly
func (view *View) Resize() {
	view.List.SetRect(0, 0, int(GetTerminalWidth()/2), int(GetTerminalHeight()/4))
}

// Render is used to update screen with new data.
func (view *View) Render(data BasicInfo) {
	view.List.Rows = []string{
		fmt.Sprintf("Core 1: [] %.1f", data.CorePercentages[0]),
		fmt.Sprintf("Core 2: [] %.1f", data.CorePercentages[1]),
		fmt.Sprintf("Core 3: [] %.1f", data.CorePercentages[2]),
		fmt.Sprintf("Core 4: [] %.1f", data.CorePercentages[3]),
		fmt.Sprintf("Memory: (%.1f/100.0)", data.MemoryUsedPercent),
		"Swap: ",
	}

	view.SystemInfoList.Rows = []string{
		fmt.Sprintf("Hostname: %s", data.Hostname),
		fmt.Sprintf("Kernel: %s", data.KernelVersion),
		fmt.Sprintf("Uptime: %s", ConvertTime(data.Uptime)),
		fmt.Sprintf("Processes: %d", data.Procs),
	}

	view.ProcessesWindow.Rows = data.Processes

	ui.Render(view.List, view.SystemInfoList, view.ProcessesWindow)
}
