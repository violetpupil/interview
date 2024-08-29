package algorithm

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(list)
	fmt.Println(list)
}
