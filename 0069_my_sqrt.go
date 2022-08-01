package main

// 速算平方根
// 执行用时：0 ms, 在所有 Go 提交中击败了 100% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 99.87% 的用户
func mySqrt(x int) int {
	var y = x
	for y*y > x {
		y = (y + x/y) / 2
	}
	return y
}
