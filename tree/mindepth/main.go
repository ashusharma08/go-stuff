package main

import "fmt"

func main() {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}
	fmt.Println("Tree Maximum Depth:", findDepth(root))
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left.Left = &TreeNode{Val: 11}
	fmt.Println("Tree Maximum Depth:", findDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		level++
		lvl := len(queue)
		for i := 0; i < lvl; i++ {
			item := queue[0]
			queue = queue[1:]
			if item.Left == nil && item.Right == nil {
				return lvl
			}
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
	}
	return level
}
