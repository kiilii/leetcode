package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 99.85% 的用户
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
