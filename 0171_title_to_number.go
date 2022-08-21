package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 64.71% 的用户
func titleToNumber(columnTitle string) int {
	var ret int
	for _, v := range columnTitle {
		ret = ret*26 + (int(v) - 66)
	}
	return ret
}
