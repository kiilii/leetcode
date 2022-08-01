package main

// 执行用时：80 ms, 在所有 Go 提交中击败了 90.83% 的用户
// 内存消耗：8.9 MB, 在所有 Go 提交中击败了 70.15% 的用户
func maxSubArray(nums []int) int {
	var ans, sum = nums[0], 0

	for _, n := range nums {
		if sum > 0 {
			sum += n
		} else {
			sum = n
		}
		if ans < sum {
			ans = sum
		}
	}
	return ans
}
