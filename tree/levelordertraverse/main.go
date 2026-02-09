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
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}
	result := traverse(root)
	fmt.Println("Level order traversal:", result)

}

func traverse(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		inte := make([]int, 0)
		loop := len(queue)
		for i := 0; i < loop; i++ {
			item := queue[0]
			queue = queue[1:]
			inte = append(inte, item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		result = append(result, inte)
		fmt.Println(len(queue))
	}
	// TODO: Write your code here
	return result
}
