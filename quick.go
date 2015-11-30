package sort

// Quick sorts data in place using the Quicksort algorithm. The function returns
// two orderings, order1 and order2, such that order1[i] is the old position of
// of the new data[i], and order2[i] is the new position of of the old data[i].
//
// https://en.wikipedia.org/wiki/Quicksort
func Quick(data []float64) ([]uint, []uint) {
	n := uint(len(data))
	order := make([]uint, 2*n)
	for i := uint(0); i < n; i++ {
		order[i] = i
	}
	if n > 1 {
		quickSort(data, order, 0, int(n)-1)
	}
	for i := uint(0); i < n; i++ {
		order[n+order[i]] = i
	}
	return order[:n], order[n:]
}

func quickSort(data []float64, order []uint, left, right int) {
	i, j := left, right
	pivot := data[(left+right)/2]
	for i <= j {
		for data[i] < pivot {
			i++
		}
		for data[j] > pivot {
			j--
		}
		if i <= j {
			data[i], data[j] = data[j], data[i]
			order[i], order[j] = order[j], order[i]
			i++
			j--
		}
	}
	if left < j {
		quickSort(data, order, left, j)
	}
	if i < right {
		quickSort(data, order, i, right)
	}
}
