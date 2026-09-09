package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testCases := [][]string{
		{"2", "1", "+", "3", "*"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		{"4", "13", "5", "/", "+"},
		{"1", "2", "+", "+"},
		{"1", "2", "3", "+"},
	}
	for _, tc := range testCases {
		fmt.Println(evaluate(tc))
	}
	fmt.Println("###########################")
	morecases := [][]string{
		{"a+b*c", "abc*+"},
		{"(a+b)*c", "ab+c*"},
		{"a+(b*(c+d))", "abcd+*+"},
		{"(a-b)+c", "ab-c+"},
		{"a^b^c", "abc^^"},
		{"a+b*c-d/e", "abc*+de/-"},
		{"a+b*(c^d-e)^(f+g*h)-i", "abcd^e-fgh*+^*+i-"},
		{"a*b^c+d", "abc^*d+"},
		{"((a+b)*c-(d-e))^(f+g)", "ab+c*de--fg+^"},
		{"a+b*(c+d*(e-f))", "abcdef-*+*+"},
		{"a+b*(c^d-e)", "abcd^e-*+"},
	}
	for idx, item := range morecases {
		out := infixToPostfix(item[0])
		if out != item[1] {
			fmt.Println(idx, ". failed expected ", item[1], "got ", out, " for ", item[0])
		}
	}
}

// func infixToPostfix(s string) string {
// 	stack := []rune{}
// 	result := []rune{}

// 	for _, ch := range s {
// 		// Operand
// 		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
// 			result = append(result, ch)
// 		} else if ch == '(' {
// 			stack = append(stack, ch)
// 		} else if ch == ')' {
// 			for len(stack) > 0 && stack[len(stack)-1] != '(' {
// 				result = append(result, stack[len(stack)-1])
// 				stack = stack[:len(stack)-1]
// 			}
// 			stack = stack[:len(stack)-1] // remove '('
// 		} else if isOperator(ch) {
// 			for len(stack) > 0 &&
// 				precedence(stack[len(stack)-1]) >= precedence(ch) {
// 				result = append(result, stack[len(stack)-1])
// 				stack = stack[:len(stack)-1]
// 			}
// 			stack = append(stack, ch)
// 		}
// 	}

// 	for len(stack) > 0 {
// 		result = append(result, stack[len(stack)-1])
// 		stack = stack[:len(stack)-1]
// 	}

// 	return string(result)
// }

func infixToPostfix(val string) string {
	var res strings.Builder
	stack := make([]rune, 0, len(val))
	precedence := map[rune]int{
		'^': 2,
		'/': 1,
		'*': 1,
		'+': 0,
		'-': 0,
	}
	for _, c := range val {
		switch c {
		case '+', '-', '*', '/', '^':
			for len(stack) > 0 && stack[len(stack)-1] != '(' &&
				(precedence[stack[len(stack)-1]] > precedence[c] ||
					(precedence[stack[len(stack)-1]] == precedence[c] && c != '^')) {
				res.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, c)
		case '(':
			stack = append(stack, c)
		case ')':
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				res.WriteRune(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		default:
			res.WriteRune(c)
		}
	}
	for len(stack) > 0 {
		res.WriteRune(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return res.String()
}

func evaluate(in []string) (int, error) {
	stack := make([]int, 0, len(in))
	for _, item := range in {
		switch item {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid data")
			}
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var res int
			switch item {
			case "*":
				res = l * r
			case "+":
				res = l + r
			case "/":
				if l == 0 {
					return 0, fmt.Errorf("cannot divide by 0 ")
				}
				res = l / r
			case "-":
				res = l - r
			}
			stack = append(stack, res)
		default:
			c, err := strconv.Atoi(item)
			if err != nil {
				return 0, err
			}
			stack = append(stack, c)
		}
	}
	if len(stack) > 1 {
		return 0, fmt.Errorf("invalid data ")
	}
	return stack[0], nil
}
