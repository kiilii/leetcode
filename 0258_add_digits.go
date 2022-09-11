package main

// 执行用时：4 ms, 在所有 Go 提交中击败了 34.96% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 100.00% 的用户
func addDigits(num int) int {
	var ret int

	for num != 0 {
		ret += num % 10
		num = num / 10
		if num == 0 && ret >= 10 {
			num = ret
			ret = 0
		}
	}
	return ret
}
