package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("abbbaaca", removeDup("abbbaaca", 3))
	fmt.Println("abbaccaa", removeDup("abbaccaa", 3))
	fmt.Println("abbacccaa", removeDup("abbacccaa", 3))

}

type stack struct {
	char  rune
	count int
}

func removeDup(str string, k int) string {
	if len(str) < k {
		return str
	}
	stk := make([]stack, 0)
	for _, item := range str {
		if len(stk) == 0 || (len(stk) > 0 && stk[len(stk)-1].char != item) {
			stk = append(stk, stack{char: item, count: 1})

		} else {

			stk[len(stk)-1].count++
		}
		for len(stk) > 0 && stk[len(stk)-1].count == k {
			stk = stk[:len(stk)-1]
		}
	}
	res := ""
	for _, it := range stk {
		res += strings.Repeat(string(it.char), it.count)
	}
	return res
}
