package main

import "fmt"

func main() {
	fmt.Println(lcs("sea", "eat"))
	fmt.Println(lcs("delete", "leet"))
}

//if str1[i] == str2[j] => do nothing. just i-1 and j-1
//else min of str1[i]+previous(i) or str2[i]+previous(j)

func lcs(str1 string, str2 string) int {
	l1, l2 := len(str1), len(str2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		dp[i][0] = int(str1[i-1]) + dp[i-1][0]
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = int(str2[i-1]) + dp[0][i-1]
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(str1[i-1]), dp[i][j-1]+int(str2[j-1]))
			}
		}
	}
	for _, item := range dp {
		fmt.Println(item)
	}
	return dp[l1][l2]
}

// (ASCII values: s=115, e=101, a=97, t=116)
// 		 s      e       a
// 		0	 115	216		313
// e	101	 0	 	0	 	0
// a	198	 0	 	0	 	0
// t	314	 0	 	0	 	0

// 		 s      e       a
// 		0	 115	216		313
// e	101	 216 	115	 	212
// a	198	 0	 	0	 	0
// t	314	 0	 	0	 	0
