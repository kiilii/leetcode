package main

import "fmt"

func Test_003() {
	// var is = [128]int{-1}

	fmt.Println(lengthOfLongestSubstring("hello world!"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
func lengthOfLongestSubstring(s string) int {
	var last = make([]int, 128)
	var start, res int
	for i := 0; i < 128; i++ {
		last[i] = -1
	}
	for index, ascii := range s {
		start = max(start, last[ascii]+1)
		res = max(res, index-start+1)
		last[ascii] = index
	}
	return res
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
