package main

import "fmt"

func Test_0013() {
	romanToInt("MCIV")
	romanToInt("LVIII")
	romanToInt("III")
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	// var signs = map[byte]int{
	// 	"I": 1,
	// 	"V": 5,
	// 	"X": 10,
	// 	"L": 50,
	// 	"C": 100,
	// 	"D": 500,
	// 	"M": 1000,
	// }
	var sum, last int
	var signs = map[uint8]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		if signs[s[i]] >= last {
			sum = sum + signs[s[i]]
		} else {
			sum = sum - signs[s[i]]
		}
		last = signs[s[i]]
	}

	fmt.Println(sum)
	return sum
}
