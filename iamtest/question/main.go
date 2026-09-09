package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(modifyString([]int{-1, 0, 1, 2, -1, -4, 4}, 0))
}
func modifyString(arr []int, k int) [][]int {
	slices.Sort(arr)
	res := make([][]int, 0)
	//-4, -1 , -1, 0, 1, 2, 4
	for idx := 0; idx < len(arr); idx++ {
		if idx > 0 && arr[idx] == arr[idx-1] {
			continue
		}

		start, end := idx+1, len(arr)-1
		diff := k - arr[idx]
		for start < end {
			if arr[start]+arr[end] == diff {
				res = append(res, []int{arr[idx], arr[start], arr[end]})
				if start < end && arr[start] == arr[start+1] {
					start++
				}
				if start < end && arr[end] == arr[end-1] {
					end--
				}
				start++
				end--
			} else if arr[start]+arr[end] > diff {
				end--
			} else {
				start++
			}
		}
	}
	return res
}
