package sortalgorithm

type ISort interface {
	Sort([]int) []int
}

type QuickSort struct {
}

func (q *QuickSort) Sort(arry []int) []int {
	var rst []int = make([]int, 10)
	return rst
}