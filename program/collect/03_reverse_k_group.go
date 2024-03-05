package program

type Node struct {
	Next  *Node
	Value int
}

// 反转单链表
func Reverse(head *Node) *Node {
	var pre *Node = nil
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

// 按分组反转单链表
func ReverseKGroup(head *Node, k int) *Node {
	first := &Node{Next: head}
	pre := first
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return first.Next
			}
		}

		next := tail.Next
		head, tail = reverseGroup(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}

	return first.Next
}

func reverseGroup(head, tail *Node) (*Node, *Node) {
	var pre *Node = nil
	cur := head
	for pre != tail {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return tail, head
}
