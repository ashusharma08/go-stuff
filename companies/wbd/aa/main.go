package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	// fmt.Println(longestNonRepeating("abvabvbb"))
	// fmt.Println(longestNonRepeating("abvbsabvbb"))

	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	// fmt.Println(searchInsertPosition([]int{1, 3, 5, 6}, 5))
	// fmt.Println(searchInsertPosition([]int{1, 3, 5, 6}, 2))

	// fmt.Println(coinChange([]int{1, 2, 5}, 11))
	// fmt.Println(coinChange([]int{1, 3, 4}, 6))

	// numCourses := 4
	// prerequisites := [][]int{
	// 	{1, 0},
	// 	{2, 0},
	// 	{3, 1},
	// 	{3, 2},
	// }

	// fmt.Println(canFinish(numCourses, prerequisites))

	// root := &TreeNode{Val: 3}

	// root.Left = &TreeNode{Val: 9}

	// root.Right = &TreeNode{Val: 20}
	// root.Right.Left = &TreeNode{Val: 15}
	// root.Right.Right = &TreeNode{Val: 7}

	// fmt.Println(levelOrder(root))
	// required := []string{"a", "b", "c"}

	// packages := []Package{
	// 	{courses: []string{"a", "b", "c"}, price: 100},
	// 	{courses: []string{"b", "x", "y"}, price: 40},
	// 	{courses: []string{"a", "y", "c"}, price: 50},
	// }

	// result := minCostPackages(required, packages)

	// fmt.Println("Minimum cost:", result)

	// fmt.Println(maxSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	// fmt.Println(temperature([]int{73, 74, 75, 71, 69, 72, 76, 73})) // [1,1,4,2,1,1,0,0]

	// points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	// k := 2

	// fmt.Println(KClosestPoints(points, k))

	fmt.Println(wordLadder([]string{"hot", "dot", "dog", "lot", "log", "cog"}, "hit", "cog"))

}

func wordLadder(list []string, begin, end string) int {
	return 0
}

func KClosestPoints(points [][]int, k int) [][]int {
	// 	17. K Closest Points to Origin
	// Given an array of points where:
	// points[i] = [x, y]
	// Return the k closest points to the origin (0,0).
	// Distance is calculated as:
	// sqrt(x² + y²)
	// Example:
	// Input:
	// points = [[1,3],[-2,2]]
	// k = 1
	// Output:
	// [[-2,2]]
	slices.SortFunc(points, func(x, y []int) int {
		d1 := (x[0] ^ 2) + (x[1] ^ 2)
		d2 := (y[0] ^ 2) + (y[1] ^ 2)

		if d1 == d2 {
			return 0
		}
		if d1 > d2 {
			return 1
		}
		return -1
	})
	return points[:k]
}

func temperature(arr []int) []int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	res := make([]int, len(arr))

	for idx := 1; idx < len(arr); idx++ {
		for len(stack) > 0 && arr[idx] > arr[stack[len(stack)-1]] {
			popIdx := stack[len(stack)-1]

			stack = stack[:len(stack)-1]
			res[popIdx] = idx - popIdx
		}
		stack = append(stack, idx)
	}

	return res
}

func maxSubarray(nums []int) int {
	currentSum := nums[0]
	mxv := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], nums[i]+currentSum)
		if currentSum > mxv {
			mxv = currentSum
		}
	}
	return mxv
}

type Package struct {
	courses []string
	price   int
}

func minCostPackages(required []string, packages []Package) int {
	type mergedPackages struct {
		covered []string
		price   int
	}
	min := math.MaxInt
	merged := make([]mergedPackages, 0)
	for i := 0; i < len(packages); i++ {

		merged = append(merged, mergedPackages{
			covered: packages[i].courses,
			price:   packages[i].price,
		})

		base := mergedPackages{
			covered: make([]string, 0),
			price:   0,
		}
		for j := i + 1; j < len(packages); j++ {
			merged = append(merged, mergedPackages{
				covered: append(packages[i].courses, packages[j].courses...),
				price:   packages[i].price + packages[j].price,
			})

			base.covered = append(base.covered, packages[j].courses...)
			base.price += packages[j].price
		}
		if len(base.covered) > 0 {
			merged = append(merged, mergedPackages{
				covered: append(packages[i].courses, base.covered...),
				price:   packages[i].price + base.price,
			})
		}
	}

	fmt.Println(merged)
	for _, v := range merged {
		mp := make(map[string]bool)
		for _, item := range v.covered {
			mp[item] = true
		}
		for _, r := range required {
			if !mp[r] {
				goto NEXT
			}
		}
		if v.price < min {
			min = v.price
		}
	NEXT:
	}
	return min
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)

	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}
		for i := 0; i < levelSize; i++ {
			x := queue[0]
			queue = queue[1:]

			level = append(level, x.Val)
			if x.Left != nil {
				queue = append(queue, x.Left)
			}
			if x.Right != nil {
				queue = append(queue, x.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
	   [[1,0],[2,1],[0,2]]
	   0: 1
	   1: 2
	   2: 0


	*/
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		a := p[0]
		b := p[1]
		adj[b] = append(adj[b], a)
		indegree[a]++
	}
	fmt.Println(adj)
	fmt.Println(indegree)
	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	fmt.Println(queue)
	count := 0

	for len(queue) > 0 {

		course := queue[0]
		queue = queue[1:]

		count++

		for _, next := range adj[course] {

			indegree[next]--

			if indegree[next] == 0 {
				queue = append(queue, next)
			}

		}

	}

	return count == numCourses
}

// https://github.com/esoptra/twintag-config/pull/158

func coinChange(coins []int, sum int) int {
	// 		0  1 2 3 4 5 6 7 8 9 10 11
	// 	1	0  1 2 3 4 5 6 7 8 9 10 11
	//  2	_  _ 1 2 2 3 3 4 4 5  5  6
	//  5   _  _ _ _ _ 1 2 2 3 3  2  3
	//

	//   0 1 2 3 4 5 6
	//1  0 1 2 3 4 5 6
	//3  0 1 2 1 2 3 2
	//4  0 1 2 3 1 2 3

	dp := make([]int, sum+1)
	i := 0
	for range sum + 1 {
		dp[i] = i
		i++
	}
	for _, coin := range coins {
		for i := coin; i < len(dp); i++ {
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
		fmt.Println(dp)
	}
	return dp[sum]
}

func searchInsertPosition(arr []int, target int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		}
	}
	return -1
}

func longestNonRepeating(str string) string {
	if len(str) <= 1 {
		return str
	}

	left, right := 0, 0
	mp := make(map[byte]int)
	start, maxLen := 0, 0

	for ; right < len(str); right++ {
		if idx, ok := mp[str[right]]; ok && idx >= left {
			left = idx + 1
		}
		mp[str[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}
	}
	return str[start : start+maxLen]
}

func groupAnagrams(arr []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, item := range arr {
		slide := [26]int{}
		for _, c := range item {
			slide[c-'a']++
		}
		mp[slide] = append(mp[slide], item)
	}
	res := make([][]string, 0)
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
