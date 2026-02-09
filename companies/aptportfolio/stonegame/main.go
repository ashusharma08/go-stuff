package main

import "fmt"

func main() {
	fmt.Println(stoneGame([]int{1, 3}, []int{2, 1}))
	fmt.Println(stoneGame([]int{1, 2}, []int{3, 1}))
	fmt.Println(stoneGame([]int{2, 4, 3}, []int{1, 6, 7}))

}

func stoneGame(arr1 []int, arr2 []int) int {
	a, b := 0, 0
	for i := 0; i < len(arr1); i++ {
		if i%2 == 0 {
			a += arr1[i]
		} else {
			b += arr2[i]
		}
	}
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}
