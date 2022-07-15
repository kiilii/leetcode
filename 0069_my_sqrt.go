package main

import "fmt"

func main() {
	fmt.Println(mySqrt(12890))
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	var y = x
	for y*y > x {
		y = (y + x/y) / 2
	}

	return y
}
