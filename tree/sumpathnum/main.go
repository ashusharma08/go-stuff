package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSumOfPathNumbers(node *TreeNode) int {
	var path int
	return findRecPath(node, path)

}

func findRecPath(node *TreeNode, path int) int {
	if node == nil {
		return 0
	}
	path = (path * 10) + node.Val

	if node.Left == nil && node.Right == nil {
		return path
	}
	return findRecPath(node.Left, path) + findRecPath(node.Right, path)
}

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Println("Total Sum of Path Numbers:", findSumOfPathNumbers(root))
}
