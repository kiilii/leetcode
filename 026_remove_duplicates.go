package main

import "fmt"

func Test_0026_remove_duplicates() {
	var nums = []int{1, 1, 2}

	fmt.Println(removeDuplicates(nums))
}

// 2023.04.30
// 执行用时：8 ms, 在所有 Go 提交中击败了 76.93% 的用户
// 内存消耗：4.2 MB, 在所有 Go 提交中击败了 64.87%的用户
func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i-1], nums[i:]...)
		}
	}
	return len(nums)
}
