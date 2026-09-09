package main

import "fmt"

func main() {
	arr := [][]int{{2, 3, 1, 1, 4}, {3, 2, 1, 0, 4}}
	for _, item := range arr {
		fmt.Println(canReachLast(item))
	}
}
func canReachLast(x []int) bool {
	// {3, 2, 1, 0, 4}
	//  F  F  F  F  X
	end := len(x) - 1

	for i := end - 1; i >= 0; i-- {
		if i+x[i] >= end {
			end = i
		}
	}

	return end == 0
}
