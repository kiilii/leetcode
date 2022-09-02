package main

// 执行用时：92 ms, 在所有 Go 提交中击败了 23.82% 的用户
// 内存消耗：7.6 MB, 在所有 Go 提交中击败了 98.02% 的用户
func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	qs(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// quick sort
func qs(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid, i := nums[0], 1
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	nums[head] = mid
	qs(nums[:head])
	qs(nums[head+1:])
}

// 执行用时：76 ms, 在所有 Go 提交中击败了 36.13% 的用户
// 内存消耗：12.3 MB, 在所有 Go 提交中击败了 5.21% 的用户
func containsDuplicate1(nums []int) bool {
	var set = make(map[int]interface{})

	for i, n := range nums {
		set[n] = struct{}{}
		if len(set) == i {
			return true
		}
	}
	return false
}
