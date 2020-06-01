package ui

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/nireo/go-manager/utils"
)

// InitUI initializes termui
func InitUI(data utils.BasicInfo) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Basic info"
	l.SetRect(0, 0, int(utils.GetTerminalWidth()), int(utils.GetTerminalHeight()/3))
	l.Rows = []string{
		fmt.Sprintf("Core 1: %.2f", data.CorePercentages[0]),
		fmt.Sprintf("Core 2: %.2f", data.CorePercentages[1]),
		fmt.Sprintf("Core 3: %.2f", data.CorePercentages[2]),
		fmt.Sprintf("Core 4: %.2f", data.CorePercentages[3]),
		"Memory: ",
		"Swap: ",
	}

	ui.Render(l)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
