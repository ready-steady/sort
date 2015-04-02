package sort

// Quick sorts an array in place using the Quicksort algorithm. The function
// returns two orderings, order1 and order2, such that order1[i] is the old
// position of of the new array[i], and order2[i] is the new position of of the
// old array[i].
//
// http://en.wikipedia.org/wiki/Quicksort
func Quick(array []float64) ([]uint, []uint) {
	n := uint(len(array))

	order := make([]uint, 2*n)

	for i := uint(0); i < n; i++ {
		order[i] = i
	}

	if n > 1 {
		quickSort(array, order, 0, int(n)-1)
	}

	for i := uint(0); i < n; i++ {
		order[n+order[i]] = i
	}

	return order[:n], order[n:]
}

func quickSort(array []float64, order []uint, left, right int) {
	i, j := left, right
	pivot := array[(left+right)/2]

	for i <= j {
		for array[i] < pivot {
			i++
		}

		for array[j] > pivot {
			j--
		}

		if i <= j {
			array[i], array[j] = array[j], array[i]
			order[i], order[j] = order[j], order[i]
			i++
			j--
		}
	}

	if left < j {
		quickSort(array, order, left, j)
	}

	if i < right {
		quickSort(array, order, i, right)
	}
}
