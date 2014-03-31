package sorter

type QuickSort struct {
}

func (q *QuickSort) Sort(arry []int) []int {
	var rst []int = make([]int, 10)

	for i := 0; i < len(rst); i++ {
		rst[i] = i
	}

	return rst
}
