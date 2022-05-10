package main

import "fmt"

func Test_0028_remove_duplicates() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("aaa", "aaaa"))
	fmt.Println(strStr("mississippi", "issipi"))
	fmt.Println(strStr("mississippi", "issi"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("mississippi", "pi"))
}

// 2023.04.30
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 77.76% 的用户
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if haystack == needle {
		return 0
	}

	var pos int
	for pos <= len(haystack)-len(needle) {
		for i := 0; i < len(needle); i++ {
			if haystack[pos+i] != needle[i] {
				pos += 1
				break
			} else if i == len(needle)-1 {
				return pos
			}
		}
	}

	if pos <= len(haystack)-len(needle) {
		return pos
	} else {
		return -1
	}
}
