package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 66.40% 的用户
func reverse(x int) int {
	var ret = 0

	for x != 0 {
		ret = (ret * 10) + x%10
		x = x / 10
	}
	if 1<<31 < ret || -1<<31 > ret {
		return 0
	}
	return ret
}
