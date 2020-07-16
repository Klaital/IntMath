package stats

import "testing"

func TestAggregator_Length(t *testing.T) {
	a := Aggregator{}
	if a.Length() != 0 {
		t.Errorf("Expected empty aggregator on init, got %d", a.Length())
		t.Fail()
	}

	a.AddItem(5)
	if a.Length() != 1 {
		t.Errorf("Expected length %d, got %d", 1, a.Length())
		t.Fail()
	}

	a.AddItem(-27)
	if a.Length() != 2 {
		t.Errorf("Expected length %d, got %d", 2, a.Length())
		t.Fail()
	}
}

func TestAggregator_Max(t *testing.T) {
	a := Aggregator{}
	a.AddItem(0)
	a.AddItem(10)
	a.AddItem(-5)

	if a.Max() != 10 {
		t.Errorf("Expected max %d, got %d", 10, a.Max())
		t.Fail()
	}

	a.AddItem(-15)
	if a.Max() != 10 {
		t.Errorf("Expected max %d, got %d", 10, a.Max())
		t.Fail()
	}

	a.AddItem(200)
	if a.Max() != 200 {
		t.Errorf("Expected max %d, got %d", 200, a.Max())
		t.Fail()
	}
}

func TestAggregator_Sum(t *testing.T) {
	a := Aggregator{}
	a.AddItem(0)
	a.AddItem(10)
	a.AddItem(-5)

	if a.Sum() != 5 {
		t.Errorf("Expected sum %d, got %d", 5, a.Max())
		t.Fail()
	}

	a.AddItem(-15)
	if a.Sum() != -10 {
		t.Errorf("Expected sum %d, got %d", -10, a.Max())
		t.Fail()
	}

	a.AddItem(200)
	if a.Sum() != 190 {
		t.Errorf("Expected sum %d, got %d", 190, a.Max())
		t.Fail()
	}
}