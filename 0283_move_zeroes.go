package main

// 执行用时：12 ms, 在所有 Go 提交中击败了 5.19% 的用户
// 内存消耗：6.4 MB, 在所有 Go 提交中击败了 98.90% 的用户
func moveZeroes(nums []int) {
	var l = len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			for j := i; j < l-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			i--
			l--
		}
	}
}
