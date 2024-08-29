package algorithm

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Log(QuickSort([]int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}))
}
