package stats

import "math"

type Aggregator struct {
	data []int64
}

func (a *Aggregator) AddItem(i int64) {
	a.data = append(a.data, i)
}
func (a *Aggregator) Length() int64 {
	return int64(len(a.data))
}
func (a *Aggregator) Sum() int64 {
	var sum int64
	for _, i := range a.data {
		sum += i
	}
	return i
}
func (a *Aggregator) Mean() int64 {
	return a.Sum() / a.Length()
}
func (a *Aggregator) MeanFloat() float64 {
	return float64(a.Sum()) / float64(a.Length())
}

func (a *Aggregator) Min() int64 {
	var min int64 = math.MaxInt64
	for _, i := range a.data {
		if i < min {
			min = i
		}
	}
	return min
}

func (a *Aggregator) Max() int64 {
	var max int64 = math.MinInt64
	for _, i := range a.data {
		if i > max {
			max = i
		}
	}
	return max
}
