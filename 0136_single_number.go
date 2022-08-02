package main

// 执行用时：8 ms, 在所有 Go 提交中击败了 99.42% 的用户
// 内存消耗：6 MB, 在所有 Go 提交中击败了 99.38% 的用户
func singleNumber(nums []int) int {
	var x = 0
	for _, v := range nums {
		x = x ^ v
	}
	return x
}
