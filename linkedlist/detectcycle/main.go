package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	orig := head
	var s, f *ListNode = head, head
	hasLoop := false
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next

		if s == f {
			hasLoop = true
			break
		}
	}
	if hasLoop {
		for orig != f {
			orig = orig.Next
			f = f.Next
		}
	}
	return f
}
func main() {
	head := &ListNode{Val: 2}
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 6}
	node3 := &ListNode{Val: 8}
	node4 := &ListNode{Val: 10}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(detectCycle(head))
}
