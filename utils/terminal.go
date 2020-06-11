package utils

import (
	ui "github.com/gizak/termui/v3"
)

// GetTerminalWidth gets the terminals width so that we can adjust the ui elements accordingly
func GetTerminalWidth() int {
	terminalWidth, _ := ui.TerminalDimensions()
	return terminalWidth
}

// GetTerminalHeight returns the height of the terminal
func GetTerminalHeight() int {
	_, terminalHeight := ui.TerminalDimensions()
	return terminalHeight
}
