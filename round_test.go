package intmath

import "testing"

func TestRoundToInt(t *testing.T) {
	y := RoundToInt(12.345)
	if y != 12 {
		t.Errorf("Failed rounding: Expected %d, got %d", 12, y)
	}

	y = RoundToInt(12.567)
	if y != 13 {
		t.Errorf("Failed rounding: Expected %d, got %d", 13, y)
	}
}

func TestRoundToPlace(t *testing.T) {
	y := RoundToPlace(12.34567, 1)
	if y != 12.30 {
		t.Errorf("Failed rounding to place: Expected %f, got %f", 12.30, y)
	}

	y = RoundToPlace(12.34567, 2)
	if y != 12.350 {
		t.Errorf("Failed rounding to place: Expected %f, got %f", 12.350, y)
	}
}
