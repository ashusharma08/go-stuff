package main

import "fmt"

func main() {
	arr1 := []int{2, 7, 9, 3, 1}
	arr2 := []int{100, 1, 1, 100}
	fmt.Println(robme(arr1))
	fmt.Println(robme(arr2))
}

func robme(arr []int) int {
	prev1, prev2 := 0, 0
	for _, item := range arr {
		temp := prev1 // we need to swap it with prev 2
		if item+prev2 > prev1 {
			prev1 = item + prev2
		}
		prev2 = temp
	}
	return prev1
}
