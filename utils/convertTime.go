package utils

import (
	"fmt"
	"strconv"
)

// ConvertTime is used to convert uptime seconds to minutes and hours
func ConvertTime(uptime uint64) string {
	hours := uptime / 3600
	minutes := (uptime % 3600) / 60
	seconds := uptime % 60

	hoursString := strconv.Itoa(int(hours))
	if hours < 10 {
		hoursString = "0" + hoursString
	}

	minutesString := strconv.Itoa(int(minutes))
	if minutes < 10 {
		minutesString = "0" + minutesString
	}

	secondsString := strconv.Itoa(int(seconds))
	if seconds < 10 {
		secondsString = "0" + secondsString
	}

	return fmt.Sprintf("%s:%s:%s", hoursString, minutesString, secondsString)
}
