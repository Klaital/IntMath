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

func TestMax(t *testing.T) {
	if x := Max(0, 1); x != 1 {
		t.Errorf("Failed to operate on (small, large). Expected %d, got %d", 1, x)
	}
	if x := Max(5, 5); x != 5 {
		t.Errorf("Failed to operate on (equal, equal). Expected %d, got %d", 5, x)
	}
	if x := Max(6, 5); x != 6 {
		t.Errorf("Failed to operate on (large, small). Expected %d, got %d", 6, x)
	}
}

func TestMaxUint64(t *testing.T) {
	if x := MaxUint64(0, 1); x != 1 {
		t.Errorf("Failed to operate on (small, large). Expected %d, got %d", 1, x)
	}
	if x := MaxUint64(5, 5); x != 5 {
		t.Errorf("Failed to operate on (equal, equal). Expected %d, got %d", 5, x)
	}
	if x := MaxUint64(6, 5); x != 6 {
		t.Errorf("Failed to operate on (large, small). Expected %d, got %d", 6, x)
	}
}
