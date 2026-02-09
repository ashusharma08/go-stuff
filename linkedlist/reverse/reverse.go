package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var p *ListNode
	var next *ListNode
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = p
		p = curr
		curr = next
	}
	return p
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
