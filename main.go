// Package sort provides sorting algorithms.
package sort

import (
	"sort"
)

// Unique sorts data in place and eliminates duplicates.
func Unique(data []float64) []float64 {
	sort.Float64s(data)
	n, k := uint(len(data)), uint(0)
	for i := uint(1); i < n; i++ {
		if data[k] != data[i] {
			k++
			data[k] = data[i]
		}
	}
	return data[:k+1]
}
