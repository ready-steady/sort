package sort

import (
	"math/rand"
	"testing"

	"github.com/ready-steady/assert"
)

func TestUnique(t *testing.T) {
	data := []float64{1, -1, 2, 0, 0, 2, 4, 1, 1, 1, 4}
	assert.Equal(Unique(data), []float64{-1, 0, 1, 2, 4}, t)
}

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique(random(1e6, 10))
	}
}

func random(count, unique uint) []float64 {
	data := make([]float64, count)
	for i := range data {
		data[i] = float64(uint(float64(unique) * rand.Float64()))
	}
	return data
}
