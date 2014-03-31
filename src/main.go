/*
main package.
*/
package main

import (
	. "./algorithms/sorter"
	"fmt"
)

func main() {
	//new
	var isort ISort = new(QuickSort)

	//init arry
	arry := make([]int, 10)
	for i := 0; i < len(arry); i++ {
		arry[i] = 10 - i
	}
	//print init arry
	fmt.Println("before sort:", arry)

	//call
	rst := isort.Sort(arry)

	//output
	fmt.Println("original list:", arry)
	fmt.Println("sorted list:", rst)

}
