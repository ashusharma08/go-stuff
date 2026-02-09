package main

import "fmt"

func main() {
	fmt.Println(valid("()[]{}"))
	fmt.Println(valid("{[}]"))
	fmt.Println(valid("(]"))
	fmt.Println(valid("{[()]}"))
}

func valid(val string) bool {
	mapMap := map[string]string{"{": "}", "[": "]", "(": ")"}
	stack := make([]string, 0)
	for _, v := range val {
		if oppVal, ok := mapMap[string(v)]; ok {
			stack = append(stack, oppVal)
		} else {
			sl := len(stack)
			if sl == 0 {
				return false
			}
			pop := stack[sl-1]
			stack = stack[:sl-1]
			if pop != string(v) {
				return false
			}
		}
	}
	return len(stack) == 0
}
