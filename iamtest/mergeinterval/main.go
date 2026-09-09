package main

import "fmt"

func main() {
	inte := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {19, 25}, {21, 24}}
	res := [][]int{inte[0]}
	i := 1
	for i < len(inte) {
		curr := res[len(res)-1]
		if curr[1] >= inte[i][0] {
			max := max(inte[i][1], curr[1])
			res[len(res)-1][1] = max
		} else {
			res = append(res, inte[i])
		}
		i++
	}
	fmt.Println(res)
}
