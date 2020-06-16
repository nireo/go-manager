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
	OS                string
	Platform          string
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
	Grid            *ui.Grid
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
	view.ProcessesWindow.RowSeparator = false

	terminalWidth, terminalHeight := ui.TerminalDimensions()
	view.Grid.SetRect(0, 0, terminalWidth, terminalHeight)
	view.Grid.Set(
		ui.NewRow(1.0/3,
			ui.NewCol(1.0/2, view.List),
			ui.NewCol(1.0/2, view.SystemInfoList),
		),
		ui.NewRow((1.0/3)*2,
			ui.NewCol(1.0, view.ProcessesWindow),
		),
	)
}

// NewView returns a pointer to a view struct
func NewView() *View {
	view := &View{
		List:            widgets.NewList(),
		SystemInfoList:  widgets.NewList(),
		ProcessesWindow: widgets.NewTable(),
		Data:            BasicInfo{},
		Grid:            ui.NewGrid(),
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
		"",
		fmt.Sprintf("[Core 1](fg:blue): %s %.1f", CPUProgressBar(data.CorePercentages[0]), data.CorePercentages[0]),
		fmt.Sprintf("[Core 2](fg:blue): %s %.1f", CPUProgressBar(data.CorePercentages[1]), data.CorePercentages[1]),
		fmt.Sprintf("[Core 3](fg:blue): %s %.1f", CPUProgressBar(data.CorePercentages[2]), data.CorePercentages[2]),
		fmt.Sprintf("[Core 4](fg:blue): %s %.1f", CPUProgressBar(data.CorePercentages[3]), data.CorePercentages[3]),
		fmt.Sprintf("[Memory](fg:blue): (%.1f/100.0)", data.MemoryUsedPercent),
		"[Swap](fg:blue): ",
	}

	view.SystemInfoList.Rows = []string{
		"",
		fmt.Sprintf("[Hostname](fg:blue): %s", data.Hostname),
		fmt.Sprintf("[Kernel](fg:blue): %s", data.KernelVersion),
		fmt.Sprintf("[Uptime](fg:blue): %s", ConvertTime(data.Uptime)),
		fmt.Sprintf("[Processes](fg:blue): %d", data.Procs),
		fmt.Sprintf("[OS](fg:blue): %s", data.OS),
		fmt.Sprintf("[Platform](fg:blue): %s", data.Platform),
	}

	view.ProcessesWindow.Rows = data.Processes

	//ui.Render(view.List, view.SystemInfoList, view.ProcessesWindow)
	ui.Render(view.Grid)
}
