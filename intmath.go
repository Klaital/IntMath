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

func Max(x int64, y int64) int64 {
	if x >= y {
		return x
	}
	return y
}

func MaxUint64(x uint64, y uint64) uint64 {
	if x >= y {
		return x
	}
	return y
}

