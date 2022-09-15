package main

// 执行用时：4 ms, 在所有 Go 提交中击败了 64.40% 的用户
// 内存消耗：2.6 MB, 在所有 Go 提交中击败了 26.85% 的用户
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var x = make(map[rune]int)

	for _, v := range s {
		x[v] += 1
	}

	for _, v := range t {
		if x[v] == 0 {
			return false
		} else {
			x[v] -= 1
		}
	}

	return true
}
