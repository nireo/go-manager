package ui

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/nireo/go-manager/utils"
)

// Run handles the ui main loop which handles starting the ui, handling event and updating the view.
func Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	view := utils.NewView()

	ev := ui.PollEvents()

	// send a tick to a time channel every 2 seconds
	tick := time.Tick(time.Second * 2)

	// main loop
	for {
		select {
		// handle different events
		case e := <-ev:
			switch e.Type {
			case ui.KeyboardEvent:
				// quit
				return
			case ui.ResizeEvent:
				view.Resize()
			}

		// handle the tick sent to time channel
		case <-tick:
			data, err := utils.LoadData()
			if err != nil {
				log.Println(err)
				break
			}

			view.Render(data)
		}
	}
}
