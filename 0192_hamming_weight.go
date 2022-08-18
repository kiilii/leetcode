package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 34.21% 的用户
func hammingWeight(num uint32) int {
	var ret uint32 = 0
	for i := 0; i < 32; i++ {
		ret += num & 1
		num >>= 1
	}
	return int(ret)
}
