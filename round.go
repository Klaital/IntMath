package intmath

import "math"

// RoundToPlace will perform rounding at the requested decimal place.
// So, RoundToPlace(12.3456, 1) => 12.30
//     RoundToPlace(12.3456, 2) => 12.350
//     RoundToPlace(12.3456, 3) => 12.3460
func RoundToPlace(x float64, decimalPlaces int64) float64 {
	multiplier := math.Pow(10.0, float64(decimalPlaces))
	tmp := math.Round(x * multiplier)
	return tmp / multiplier
}

// RoundToInt will round a float, and return as an integer
func RoundToInt(x float64) int64 {
	return int64(math.Round(x))
}
