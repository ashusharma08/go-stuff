package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(replaceDigits("a1b2c3d4e"))
	fmt.Println(maxSubArray([]int{-1}))
}

func sortSentence(s string) string {
	values := strings.Split(s, " ")

	res := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		l := len(values[i])
		num := values[i][l-1:]
		n, _ := strconv.ParseInt(num, 10, 0)
		res[n-1] = values[i][:l-1]
	}

	return strings.Join(res, " ")
}
func replaceDigits(s string) string {
	runes := []rune(s)
	for i := 1; i < len(runes); i += 2 {
		v := runes[i]
		if i+1 >= len(runes) {
			continue
		}
		n := runes[i+1]
		parse, _ := strconv.ParseInt(string(n), 10, 0)
		runes[i+1] = rune(int64(v) + parse)
	}
	return string(runes)
}
func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	currentSum := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = maxint(nums[i], nums[i]+currentSum)
		if currentSum > max {
			max = currentSum
		}
	}
	return max
}

// 0 0 0
// 0 1 0
// 1 1 1

// 0 0 1
// 0 1 1
// 0 0 1

// 0 0 0
// 0 1 0
// 1 1 1

// for

// 1 1 1
// 0 1 0
// 0 0 0

// r := [4]bool{true, true, true, true}
// 	n := len(mat)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			// 旋转 0 度
// 			if mat[i][j] != target[i][j] {
// 				r[0] = false
// 			}
// 			// 旋转 90 度
// 			if mat[i][j] != target[j][n-i-1] {
// 				r[1] = false
// 			}
// 			// 旋转 180 度
// 			if mat[i][j] != target[n-i-1][n-j-1] {
// 				r[2] = false
// 			}
// 			// 旋转 270 度
// 			if mat[i][j] != target[n-j-1][i] {
// 				r[3] = false
// 			}
// 		}
// 	}
// 	// 只要其中一个角度匹配，即表示OK
// 	return r[0] || r[1] || r[2] || r[3]
