package qsort

import (
	"testing"
)

var isort ISort

func Init() {
	isort = new(QuickSort)
}

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	//call
	isort.Sort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("QuickSort failed. Expected 1 2 3 4 5,Actual ", values)
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	//call
	isort.Sort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("QuickSort failed. Expected 1 2 3 5 5,Actual ", values)
	}
}

func TestQuickSort3(t *testing.T) {
	values := []int{5}
	//call
	isort.Sort(values)

	if values[0] != 5 {
		t.Error("QuickSort failed. Expected 5,Actual ", values)
	}
}
