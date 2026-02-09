package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}

func buildArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] + len(nums)*(nums[nums[i]]%len(nums))
	}
	for i, v := range nums {
		nums[i] = v / len(nums)
	}
	return nums
}

// 0 2 1 5 3 4
//
