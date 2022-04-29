package main

import (
	"fmt"
)

func Test_009() {
	var i int = -121

	fmt.Println(isPalindrome(i))
}
func isPalindrome(x int) bool {
	var tmp = x
	var y int

	if x < 0 {
		return false
	}

	for tmp != 0 {
		y = (y * 10) + (tmp % 10)
		tmp = tmp / 10
	}

	return x == y
}
