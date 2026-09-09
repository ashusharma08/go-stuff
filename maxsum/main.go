package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSum([]int{-3, -2, -5}))
	fmt.Println(maxSum([]int{5, -3, 5}))

}

func maxSum(arr []int) int {
	currSum := arr[0]
	maxSum := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]+currSum < arr[i] {
			currSum = arr[i]
		} else {
			currSum += arr[i]
		}
		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
