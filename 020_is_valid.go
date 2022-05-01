package main

import "fmt"

func Test_020_isValid() {
	var cases = map[string]bool{
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
		"({[)":   false,
	}
	for c, res := range cases {
		if got := isValid(c); got != res {
			fmt.Printf("isValid() = %v, want %v", got, res)
		}
	}
}

// 2023.04.29
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了 9.19% 的用户
func isValid(s string) bool {
	// var left = []string{"(", "[", "{"}
	// var right = []string{")", "]", "}"}
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
	var stack = make([]string, 0)

	for _, x := range s {
		if signs[string(x)] {
			stack = append(stack, string(x))
		} else {
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] != brackets[string(x)] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
