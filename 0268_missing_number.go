package main

// 执行用时：12 ms, 在所有 Go 提交中击败了 94.42% 的用户
// 内存消耗：6 MB, 在所有 Go 提交中击败了 72.36% 的用户
func missingNumber(nums []int) int {
	var ret = len(nums)

	for i := range nums {
		ret ^= nums[i]
		ret ^= i
	}

	return ret
}
