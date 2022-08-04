package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3 MB, 在所有 Go 提交中击败了 35.12% 的用户
func isPalindrome0125(s string) bool {
	var x = []byte{}
	for i := 0; i < len(s); i++ {
		if (s[i]&0xDF >= 65 && s[i]&0xDF <= 90) || (s[i] <= 57 && s[i] >= 48) {
			x = append(x, byte(s[i]&0xDF))
		}
	}

	var i, j = 0, len(x) - 1
	for i <= j {
		if x[i] != x[j] {
			return false
		}
		i++
		j--
	}

	return true
}
