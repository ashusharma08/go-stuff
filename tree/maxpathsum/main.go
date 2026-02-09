package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Solution struct {
	globalMaximumSum int
}

// NewSolution creates a new instance of Solution
func NewSolution() *Solution {
	return &Solution{globalMaximumSum: math.MinInt32}
}

// findMaximumPathSum starts the recursive process and returns the result
func (s *Solution) findMaximumPathSum(root *TreeNode) int {
	s.globalMaximumSum = math.MinInt32 // Reset the global maximum sum for each new tree
	s.findMaximumPathSumRecursive(root)
	return s.globalMaximumSum
}
func (s *Solution) findMaximumPathSumRecursive(currentNode *TreeNode) int {
	if currentNode == nil {
		return 0
	}
	leftSum := s.findMaximumPathSumRecursive(currentNode.Left)
	rightSum := s.findMaximumPathSumRecursive(currentNode.Right)

	leftSum = max(leftSum, 0)
	rightSum = max(rightSum, 0)
	if leftSum+rightSum+currentNode.Val > s.globalMaximumSum {
		s.globalMaximumSum = leftSum + rightSum + currentNode.Val
	}
	return max(leftSum, rightSum) + currentNode.Val
}

func main() {
	sol := NewSolution()
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println("Maximum Path Sum:", sol.findMaximumPathSum(root))

	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Right.Left.Left = &TreeNode{Val: 7}
	root.Right.Left.Right = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 9}
	fmt.Println("Maximum Path Sum:", sol.findMaximumPathSum(root))

	root = &TreeNode{Val: -1}
	root.Left = &TreeNode{Val: -3}
	fmt.Println("Maximum Path Sum:", sol.findMaximumPathSum(root))
}
