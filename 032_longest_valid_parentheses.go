package main

import "fmt"

var signs = map[string]bool{
	"(": true,
	"[": true,
	"{": true,
	")": false,
	"}": false,
	"]": false,
}
var brackets = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func Test_0032() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(")()(())"))
	fmt.Println(longestValidParentheses(")())())"))
	fmt.Println(longestValidParentheses("()(()"))
}

func longestValidParentheses(ss string) int {
	var stack = make([]string, 0)
	var num, max int
	for _, s := range ss {
		if signs[string(s)] {
			stack = append(stack, string(s))
		} else {
			if len(stack) <= 0 {
				num = 0
				continue
			}
			if stack[len(stack)-1] == brackets[string(s)] {
				stack = stack[:len(stack)-1]
				num += 1
			} else {
				num = 0
			}
			if num > max {
				max = num
			}
		}
	}

	return max * 2
}
