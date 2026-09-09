package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next = &ListNode{Val: 10}

	result := reverse(head)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}

func reverse(l *ListNode) *ListNode {
	var prev, curr, next *ListNode = nil, l, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
