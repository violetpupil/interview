package program

type Node struct {
	Next  *Node
	Value int
}

// 按分组反转单链表
func ReverseKGroup(head *Node, k int) {}

// 反转单链表
func Reverse(head *Node) *Node {
	var pre *Node = nil
	cur := head
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}
