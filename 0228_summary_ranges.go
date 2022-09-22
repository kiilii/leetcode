package main

import (
	"fmt"
	"strings"
)

// 执行用时：0 ms, 在所有 Go 提交中击败了 100% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 76.36% 的用户
func summaryRanges(nums []int) (ret []string) {
	var tmp = make([]string, 0, 2)
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			if len(tmp) == 0 {
				tmp = append(tmp, fmt.Sprint(nums[i]))
			}
		} else {
			tmp = append(tmp, fmt.Sprint(nums[i]))
			ret = append(ret, strings.Join(tmp, "->"))
			tmp = tmp[:0]
		}
	}
	return ret
}
