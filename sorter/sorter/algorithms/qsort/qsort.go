package qsort

type QuickSort struct {
}

func (q *QuickSort) Sort(array []int) {
	var length = len(array)
	qsort(array, 0, length-1)
}

func find_pivot(array []int, l, r int) int {

	var i, j, key = l, r, array[l]

	for {
		if j > i && array[j] >= key {
			j--
		}
		if i == j {
			return i
		}
		//exchange
		array[j], array[i] = array[i], array[j]

		for {
			if i < j && array[i] <= key {
				i++
			}
			if i == j {
				return i
			}
			array[i], array[j] = array[j], array[i]
			break
		}

	}
	return i
}

func qsort(array []int, l, r int) {
	var pivot = find_pivot(array, l, r)
	if pivot-1 > l {
		qsort(array, l, pivot-1)
	}
	if pivot+1 < r {
		qsort(array, pivot+1, r)
	}
}
