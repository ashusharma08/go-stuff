package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution struct definition
type Solution struct{}

// FindSuccessor finds the next node in the tree after the node with the given key
func (s *Solution) FindSuccessor(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	var prev *TreeNode
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lvl := len(queue)
		for i := 0; i < lvl; i++ {
			item := queue[0]
			queue = queue[1:]
			if prev != nil && prev.Val == key {
				return item
			}
			prev = item
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
	}

	return nil
}

func main() {
	s := Solution{}

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	result := s.FindSuccessor(root, 3)
	if result != nil {
		fmt.Println(result.Val, " ")
	}

	root = &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}

	result = s.FindSuccessor(root, 9)
	if result != nil {
		fmt.Println(result.Val, " ")
	}

	result = s.FindSuccessor(root, 12)
	if result != nil {
		fmt.Println(result.Val, " ")
	}
}
