package main

// 执行用时：12 ms, 在所有 Go 提交中击败了 94.24% 的用户
// 内存消耗：5.9 MB, 在所有 Go 提交中击败了 76.03% 的用户
func majorityElement(nums []int) int {
	var ret, count int
	for _, v := range nums {
		if count == 0 {
			ret = v
			count++
		} else if ret == v {
			count++
		} else {
			count--
		}
	}
	return ret
}

// 执行用时：16 ms, 在所有 Go 提交中击败了 41.49% 的用户
// 内存消耗：5.9 MB, 在所有 Go 提交中击败了 15.21% 的用户
func majorityElement2(nums []int) int {
	var x = make(map[int]int, 0)
	var ret, count int
	for _, v := range nums {
		x[v] += 1
		if x[v] > count {
			count = x[v]
			ret = v
		}
	}
	return ret
}
