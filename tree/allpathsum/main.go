package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}
	sum := 23
	result := findPaths(root, sum)
	fmt.Printf("Tree paths with sum %d: %v\n", sum, result)
}
func findPaths(root *TreeNode, sum int) [][]int {
	allPaths := make([][]int, 0)
	var currPath []int
	findRecPaths(root, currPath, sum, &allPaths)
	return allPaths
}

func findRecPaths(node *TreeNode, currPath []int, sum int, allPaths *[][]int) {
	if node == nil {
		return
	}
	currPath = append(currPath, node.Val)

	if node.Val == sum && node.Left == nil && node.Right == nil {
		path := make([]int, len(currPath))
		copy(path, currPath)
		*allPaths = append(*allPaths, path)
	} else {
		findRecPaths(node.Left, currPath, sum-node.Val, allPaths)
		findRecPaths(node.Right, currPath, sum-node.Val, allPaths)

	}
	currPath = currPath[:len(currPath)-1]

}
