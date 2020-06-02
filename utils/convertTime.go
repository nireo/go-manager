package utils

import "fmt"

// ConvertTime is used to convert uptime seconds to minutes and hours
func ConvertTime(uptime uint64) string {
	hours := uptime / 3600
	minutes := (uptime % 3600) / 60
	seconds := uptime % 60
	return fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)
}
