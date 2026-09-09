package main

import (
	"fmt"
)

func main() {
	bst := &BST{}
	bst.val = 10
	bst.left = &BST{
		val: 5,
	}
	bst.right = &BST{
		val: 15,
		left: &BST{
			val: 6,
		},
		right: &BST{
			val: 20,
		},
	}
	fmt.Println(isValid(bst))
}

type BST struct {
	val   int
	left  *BST
	right *BST
}

func isValid(b *BST) bool {
	st := make([]*BST, 0)
	curr := b
	var prev *int
	for curr != nil && len(st) > 0 {
		//go all the way left
		for curr != nil {
			st = append(st, curr)
			curr = curr.left
		}
		curr := st[len(st)-1]
		st = st[:len(st)-1]

		if prev != nil && *prev < curr.val {
			return false
		}
		prev = &curr.val
		curr = curr.right

	}
	return true
}
