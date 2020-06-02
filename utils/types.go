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
}

// View data structure
type View struct {
	List *widgets.List
	Data BasicInfo
}

// Init initializes the ui
func (view *View) Init() {
	view.List.Title = "Basic information"
	view.List.Border = false
	view.List.SetRect(0, 0, int(GetTerminalWidth()), int(GetTerminalHeight()))
}

// NewView returns a pointer to an view struct
func NewView() *View {
	view := &View{
		List: widgets.NewList(),
		Data: BasicInfo{},
	}

	view.Init()
	return view
}

// Resize is used when an terminal resize event occurs and it updates list dimensions accordinly
func (view *View) Resize() {
	view.List.SetRect(0, 0, int(GetTerminalWidth()), int(GetTerminalHeight()))
}

// Render is used to update screen with new data.
func (view *View) Render(data BasicInfo) {
	view.List.Rows = []string{
		fmt.Sprintf("Core 1: %.2f", data.CorePercentages[0]),
		fmt.Sprintf("Core 2: %.2f", data.CorePercentages[1]),
		fmt.Sprintf("Core 3: %.2f", data.CorePercentages[2]),
		fmt.Sprintf("Core 4: %.2f", data.CorePercentages[3]),
		"Memory: ",
		"Swap: ",
	}

	ui.Render(view.List)
}
