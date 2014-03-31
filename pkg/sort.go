package sortalgorithm

import (
	"fmt"
)

type ISort interface {
	Sort([]int) []int
}

type QuickSort struct {
}

func (q *QuickSort) Sort(arry []int) []int {
	var rst []int

	copy(rst, arry)

	fmt.Println("in sort:", arry)

	return rst
}
