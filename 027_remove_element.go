package main

import "fmt"

func Test_0027_remove_element() {
	var nums = []int{1, 1, 2}

	fmt.Println(removeElement(nums, 1))
}

// 2023.04.30
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 61.13% 的用户
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
	}
	return len(nums)
}
