package main

// 执行用时：28 ms, 在所有 Go 提交中击败了 87.24% 的用户
// 内存消耗：6.4 MB, 在所有 Go 提交中击败了 60.06% 的用户
func reverseString(s []byte) {
	var l = len(s) - 1
	for i := range s {
		if i < l-i {
			s[i], s[l-i] = s[l-i], s[i]
		} else {
			return
		}
	}
}
