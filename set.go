package intmath

type Set struct {
	data map[int64]bool
}

func NewSet() *Set {
	return &Set{
		data: make(map[int64]bool, 0),
	}
}

func (s *Set) GetItems() []int64 {
	uniques := make([]int64, 0, s.Length())
	for datum, present := range s.data {
		if present {
			uniques = append(uniques, datum)
		}
	}
	return uniques
}

func (s *Set) Add(item int64) {
	s.data[item] = true
}

func (s *Set) Delete(item int64) {
	s.data[item] = false
}

func (s *Set) Check(item int64) bool {
	return s.data[item]
}

func (s *Set) Length() int64 {
	var count int64
	for _, present := range s.data {
		if present {
			count ++
		}
	}
	return count
}