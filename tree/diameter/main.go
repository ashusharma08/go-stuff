package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	treeDiameter int
}

func (s *Solution) findDiameter(root *TreeNode) int {
	s.treeDiameter = 0
	s.calculateHeight(root)
	return s.treeDiameter
}

func (s *Solution) calculateHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := s.calculateHeight(root.Left)
	rightHeight := s.calculateHeight(root.Right)
	if leftHeight+rightHeight+1 > s.treeDiameter {
		s.treeDiameter = leftHeight + rightHeight + 1
	}
	return max(leftHeight, rightHeight) + 1
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	sol := Solution{}
	fmt.Println("Tree Diameter:", sol.findDiameter(root))

	root.Left.Left = nil
	root.Right.Left.Left = &TreeNode{Val: 7}
	root.Right.Left.Right = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 9}
	root.Right.Left.Right.Left = &TreeNode{Val: 10}
	root.Right.Right.Left.Left = &TreeNode{Val: 11}

	fmt.Println("Tree Diameter:", sol.findDiameter(root))
}
