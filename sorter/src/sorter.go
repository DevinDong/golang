package main

import (
	"algorithms/qsort"
	"fmt"
	"time"
)

func main() {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var sorter = new(qsort.QuickSort)
	sorter.Sort(array)
	fmt.Println(array)
}
