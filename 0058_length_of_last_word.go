package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 99.54% 的用户
func lengthOfLastWord(s string) int {
	var ans = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 && ans == 0 {
			continue
		}

		if s[i] == 32 && ans != 0 {
			return ans
		} else {
			ans++
		}
	}

	return ans
}
