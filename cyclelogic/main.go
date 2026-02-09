package main

import "fmt"

func main() {
	// fmt.Println(findNumbers([]int{2, 3, 1, 8, 2, 3, 5, 1}))
	fmt.Println(MinCoins([]int{1, 2, 4}, 11))
}

func findNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	res := make([]int, 0)
	for i, val := range nums {
		if i+1 != val {
			res = append(res, i+1)
		}
	}

	return res
}

func MinCoins(coins []int, sum int) int {
	dp := make([]int, sum+1)
	i := 0
	for range sum + 1 {
		dp[i] = sum
		i++
	}
	dp[0] = 0
	fmt.Println("____", dp)

	for _, c := range coins {
		for i := c; i < len(dp); i++ {
			dp[i] = min(dp[i], dp[i-c]+1)
		}
		fmt.Println(c, "__", dp)
	}
	if dp[sum] > sum {
		return -1
	}
	return dp[sum]
}

// 0 1 2 3 4 5 6

// 0 6 6 6 6 6 6
// 0 1 2 3 4 5 6
// 0 1
