package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(maxArea([]int{2, 0, 2}))

	fmt.Println(maxArea([]int{2, 4}))
	fmt.Println(maxArea([]int{2, 1, 2}))

}

func maxArea(v []int) int {
	mx := 0
	stack := make([]int, 0, len(v))
	v = append(v, 0)

	for i := 0; i < len(v); i++ {
		item := v[i]
		for len(stack) > 0 && item < v[stack[len(stack)-1]] {
			top := v[stack[len(stack)-1]] // get the value from v for the top index of stack
			stack = stack[:len(stack)-1]  //stack = stack top removed
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			area := width * top
			if area > mx {
				mx = area
			}
		}
		stack = append(stack, i)
	}

	return mx
}

func pop(arr []int) ([]int, int) {
	return arr[1:], arr[0]
}

/*
   2,1,2
   2 area 2
   1,2
   2*1
   1*2

*/
