package main

import (
	"fmt"
	"math"
)

func Test_0014() {
	var strs = []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

// 2023.04.29
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了 62.11%的用户
func longestCommonPrefix(strs []string) string {
	var min = math.MaxInt32

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) <= min {
			min = len(strs[i])
		}
	}

	var i = 0
	for ; i < min; i++ {
		var flag = strs[0][i]
		for _, str := range strs {
			if str[i] != flag {
				return string(strs[0][:i])
			}
		}
	}
	return string(strs[0][:i])
}
