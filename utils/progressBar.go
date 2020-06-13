package utils

// CPUProgressBar returns a simple progress bar used with displaying cpu percentages
func CPUProgressBar(percentage float64) string {
	var barAmount int = int(percentage / 5)
	endString := ""

	barsInString := 0
	for {
		if barsInString == barAmount {
			break
		}
		endString += "|"
		barsInString++
	}

	return endString
}
