package intmath

import "math"

func Abs(x int64) int64 {
	if x > 0 {
		return x
	}

	if x == math.MinInt64 {
		return math.MaxInt64
	}

	return -x
}
