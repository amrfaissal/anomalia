package anomalia

import (
	"sort"
)

type mapper func(float64) float64
type mapperWithIndex func(int, float64) float64
type predicate func(float64) bool

func minMax(data []float64) (float64, float64) {
	var (
		max float64 = data[0]
		min float64 = data[0]
	)
	for _, value := range data {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func mapSlice(slice []float64, mapper mapper) []float64 {
	for idx, value := range slice {
		slice[idx] = mapper(value)
	}
	return slice
}

func mapSliceWithIndex(slice []float64, mapper mapperWithIndex) []float64 {
	for idx, value := range slice {
		slice[idx] = mapper(idx, value)
	}
	return slice
}

func filter(slice []float64, predicate predicate) (ret []float64) {
	for _, value := range slice {
		if predicate(value) {
			ret = append(ret, value)
		}
	}
	return
}

func copySlice(input []float64) []float64 {
	s := make([]float64, len(input))
	copy(s, input)
	return s
}

func sortedCopy(input []float64) (copy []float64) {
	copy = copySlice(input)
	sort.Float64s(copy)
	return
}

func insertAt(slice []float64, pos int, elem float64) []float64 {
	if pos < 0 {
		pos = 0
	} else if pos >= len(slice) {
		pos = len(slice)
	}
	out := make([]float64, len(slice)+1)
	copy(out[:pos], slice[:pos])
	out[pos] = elem
	copy(out[pos+1:], slice[pos:])
	return out
}
