package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 69.06% 的用户
func isHappyV1(n int) bool {
	var x = 0
	var logs = make(map[int]bool, 0)

	for x != 1 {
		x = 0
		for n > 0 {
			x += (n % 10) * (n % 10)
			n /= 10
		}
		if _, has := logs[x]; has {
			return false
		} else {
			logs[x] = true
		}
		n = x
	}

	return true
}
