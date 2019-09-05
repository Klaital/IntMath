package intmath

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	if x := Abs(1); x != 1 {
		t.Errorf("Failed to operate on positive int. Expected 1, got %d", x)
	}

	if x := Abs(-1); x != 1 {
		t.Errorf("Failed to operate on a negative int. Expected 1, got %d", x)
	}

	if x := Abs(math.MaxInt64); x != math.MaxInt64 {
		t.Errorf("Failed to operate on MaxInt64. Expected %d, got %d", math.MaxInt64, x)
	}

	if x := Abs(math.MinInt64); x != math.MaxInt64 {
		t.Errorf("Failed to operate on MinInt64. Expected MaxInt (%d), got %d", math.MaxInt64, x)
	}
}
