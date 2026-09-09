package main

import "fmt"

func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Print(findGap(temp))
}

func findGap(arr []int) []int {
	/*
		0(73)

	*/

	stack := make([]int, 0)
	res := make([]int, len(arr))
	for idx, item := range arr {
		for len(stack) > 0 && item > arr[stack[len(stack)-1]] {
			pr := stack[len(stack)-1]
			res[pr] = idx - pr
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idx)
	}
	return res
}
