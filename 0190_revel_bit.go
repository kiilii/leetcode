package main

// 执行用时：4 ms, 在所有 Go 提交中击败了 12.97% 的用户
// 内存消耗：2.4 MB, 在所有 Go 提交中击败了 98.02% 的用户
func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		ret <<= 1
		ret += num & 1
		num >>= 1
	}
	return ret
}
