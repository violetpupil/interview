package program

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := Reverse(n1)

	for ; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}
