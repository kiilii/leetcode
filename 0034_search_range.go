package main

// 在排序数组中查找元素的第一个和最后一个位置
// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]

// 执行用时：4 ms, 在所有 Go 提交中击败了 96.46% 的用户
// 内存消耗：3.8 MB, 在所有 Go 提交中击败了 57.51% 的用户
func searchRange(nums []int, target int) []int {
	var ret = []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	var left, right = 0, len(nums) - 1
	var mid = (left + right) / 2
	// 左边界
	for left <= right {
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			ret[0] = mid
			right = mid - 1
		}
		mid = (left + right) / 2
	}

	// 右边界
	left, right = 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			ret[1] = mid
			left = mid + 1
		}
	}

	return ret
}
