package main

// 搜索旋转排序数组

// 二次提交分别

// 执行用时：4 ms, 在所有 Go 提交中击败了 11.90% 的用户
// 内存消耗：2.4 MB, 在所有 Go 提交中击败了 100.00% 的用户

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.4 MB, 在所有 Go 提交中击败了 58.66% 的用户
func search(nums []int, target int) int {
	var head, tail = 0, len(nums) - 1
	if nums[head] == target {
		return head
	}
	if nums[tail] == target {
		return tail
	}

	for head <= tail {
		var mid = (head + tail) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[tail] { // 右边有序递增
			if target <= nums[tail] && nums[mid] < target {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		} else { // 左边有序递增
			if target >= nums[head] && nums[mid] > target {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		}
	}

	return -1
}
