package main

// 执行用时：92 ms, 在所有 Go 提交中击败了 73.34% 的用户
// 内存消耗：10.9 MB, 在所有 Go 提交中击败了 41.87% 的用户
func containsNearbyDuplicate(nums []int, k int) bool {
	var set = make(map[int]int)

	for i, n := range nums {
		if ii, has := set[n]; !has {
			set[n] = i
		} else {
			if i-ii <= k {
				return true
			} else {
				set[n] = i
			}
		}
	}
	return false
}
