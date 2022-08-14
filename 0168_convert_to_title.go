package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 81.09% 的用户
func convertToTitle1(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	var s []rune
	for columnNumber > 0 {
		columnNumber -= 1
		s = append(s, rune(columnNumber%26+65))
		columnNumber /= 26
	}

	var i, j = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return string(s)
}

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 77.72% 的用户
func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	var s []rune
	for columnNumber > 0 {
		columnNumber -= 1
		s = append(s, rune(columnNumber%26+65))
		columnNumber /= 26
	}

	var ret string
	for i := len(s) - 1; i >= 0; i-- {
		ret += string(s[i])
	}

	return ret
}
