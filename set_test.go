package intmath

import "testing"

func TestSet(t *testing.T) {
	s := NewSet()

	if s.Length() != 0 {
		t.Errorf("Initialized set should have length 0. Got %d", s.Length())
		t.Fail()
	}

	// Add then delete the same item
	s.Add(5)
	if s.Length() != 1 {
		t.Errorf("Set should have length %d. Got %d", 1, s.Length())
		t.Fail()
	}
	s.Delete(5)
	if s.Length() != 0 {
		t.Errorf("Set should have length %d. Got %d", 0, s.Length())
		t.Fail()
	}

	// Add an item, then delete something that doesn't exist
	s.Add(5)
	if s.Length() != 1 {
		t.Errorf("Set should have length %d. Got %d", 1, s.Length())
		t.Fail()
	}
	s.Delete(1)
	if s.Length() != 1 {
		t.Errorf("Set should have length %d. Got %d", 1, s.Length())
		t.Fail()
	}

	// Add multiple items, then delete some of them
	s.Add(5)
	if s.Length() != 1 {
		t.Errorf("Set should have length %d. Got %d", 1, s.Length())
		t.Fail()
	}
	s.Add(1)
	if s.Length() != 2 {
		t.Errorf("Set should have length %d. Got %d", 2, s.Length())
		t.Fail()
	}
	s.Add(3)
	if s.Length() != 3 {
		t.Errorf("Set should have length %d. Got %d", 3, s.Length())
		t.Fail()
	}
	s.Delete(1)
	if s.Length() != 2 {
		t.Errorf("Set should have length %d. Got %d", 2, s.Length())
		t.Fail()
	}
	s.Delete(3)
	if s.Length() != 1 {
		t.Errorf("Set should have length %d. Got %d", 1, s.Length())
		t.Fail()
	}

	if s.Check(1) {
		t.Errorf("Expected item %d to be missing, but it is present.", 1)
		t.Fail()
	}
	if !s.Check(5) {
		t.Errorf("Expected item %d to be present, but it is missing.", 5)
		t.Fail()
	}
}
