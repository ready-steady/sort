package sort

import (
	"testing"

	"github.com/ready-steady/support/assert"
)

func TestQuickSort(t *testing.T) {
	data := []float64{3, 1, 6, 9, 4, 2}

	order1, order2 := Quick(data)

	assert.Equal(data, []float64{1, 2, 3, 4, 6, 9}, t)
	assert.Equal(order1, []uint{1, 5, 0, 4, 2, 3}, t)
	assert.Equal(order2, []uint{2, 0, 4, 5, 3, 1}, t)
}
