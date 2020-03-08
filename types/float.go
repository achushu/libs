package types

import "math"

// RoundFloat rounds x to the nearest given unit
// ex: given x=1.1234, unit=0.001 returns 1.123
func RoundFloat(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
