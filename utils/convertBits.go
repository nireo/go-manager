package utils

import "math"

// ConvertBits converts an amount of bits to gigabytes
func ConvertBits(bits uint64) float64 {
	return math.Pow10(-9) * float64(bits)
}
