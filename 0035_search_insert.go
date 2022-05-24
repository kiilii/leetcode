package main

import "fmt"

func Test_0035_search_insert() {
	var nums = []int{1, 3, 5, 6, 7}
	nums = []int{1}

	fmt.Println(searchInsert(nums, 1))
}

// 2023.05.22
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了 57.56% 的用户
func searchInsert(nums []int, target int) int {
	var left, right, mid int
	right = len(nums)
	mid = (left + len(nums)) / 2

	for {
		if left == mid || right == mid {
			if nums[mid] >= target {
				return mid
			} else {
				return mid + 1
			}
		}

		if nums[mid] < target {
			left = mid
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid
		}

		mid = (left + right) / 2
	}
}
