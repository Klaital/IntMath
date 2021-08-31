package intmath

import "math"

// Abs computes the absolute value of x.
// Exception: MinInt turns into MaxInt
func Abs(x int64) int64 {
	if x > 0 {
		return x
	}

	if x == math.MinInt64 {
		return math.MaxInt64
	}

	return -x
}

// Max returns the greater of the two input parameters.
func Max(x int64, y int64) int64 {
	if x >= y {
		return x
	}
	return y
}

// MaxUint64 returns the greater of the two input parameters.
func MaxUint64(x uint64, y uint64) uint64 {
	if x >= y {
		return x
	}
	return y
}

