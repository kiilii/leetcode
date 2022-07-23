package main

// 执行用时：68 ms, 在所有 Go 提交中击败了 66.67% 的用户
// 内存消耗：8.3 MB, 在所有 Go 提交中击败了 87.82% 的用户
func maxArea2(height []int) int {
	var ret, head int
	var tail = len(height) - 1
	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for head < tail {
		if height[head] < height[tail] {
			ret = max(ret, (tail-head)*height[head])
			head++
		} else {
			ret = max(ret, (tail-head)*height[tail])
			tail--
		}
	}
	return ret
}

// 执行用时：80 ms, 在所有 Go 提交中击败了 15.22% 的用户
// 内存消耗：8.5 MB, 在所有 Go 提交中击败了 48.97% 的用户
func maxArea1(height []int) int {
	var ret, head int
	var tail = len(height) - 1

	for head < tail {
		if height[head] < height[tail] {
			if (tail-head)*height[head] > ret {
				ret = (tail - head) * height[head]
			}
			head++
		} else {
			if (tail-head)*height[tail] > ret {
				ret = (tail - head) * height[tail]
			}
			tail--
		}
	}

	return ret
}
