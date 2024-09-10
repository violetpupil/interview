package program

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	list := []int{2, 3, 1, 0, 2, 5, 3}
	t.Log(FindDuplicate(list))
}

func TestFindDuplicateSpace(t *testing.T) {
	list := []int{2, 3, 5, 4, 3, 2, 6, 7}
	t.Log(FindDuplicateSpace(list))
}
