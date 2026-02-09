package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	result := reverse(head, 2, 4)
	fmt.Print("Nodes of the reversed LinkedList are: ")
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}

func reverse(head *ListNode, p int, q int) *ListNode {
	var current, prev *ListNode = head, nil
	for i := 0; current != nil && i < p-1; i++ {
		prev = current
		current = current.Next
	}

	nodeBeforeP := prev
	nodeAtP := current

	var next *ListNode

	for i := 0; current != nil && i < q-p+1; i++ {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	if nodeBeforeP != nil {
		nodeBeforeP.Next = prev
	} else {
		head = prev
	}
	nodeAtP.Next = current
	return head
}
