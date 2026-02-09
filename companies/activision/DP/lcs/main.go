package main

import "fmt"

func main() {
	fmt.Println(lcs("leet", "delete"))
}

func lcs(str1 string, str2 string) int {
	l1, l2 := len(str1), len(str2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1][l2]
}
