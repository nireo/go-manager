package utils

// CPUProgressBar returns a simple progress bar used with displaying cpu percentages
func CPUProgressBar(percentage float64) string {
	var barAmount int = int(percentage / 5)
	startSring := "||||||||||||||||||||"
	if barAmount == 0 {
		return startSring
	}
	endString := "[" + startSring[0:barAmount] + "](fg:red)" + startSring[barAmount-1:]

	return endString
}
