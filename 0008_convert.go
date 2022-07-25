package main

// 执行用时：16 ms, 在所有 Go 提交中击败了 28.63% 的用户
// 内存消耗：6.9  MB, 在所有 Go 提交中击败了 29.22% 的用户
func convert(s string, numRows int) string {
	var seg = make([]string, numRows)
	var index int

	if numRows <= 1 {
		return s
	}

	for i, v := range s {
		index = i % (numRows - 1)
		if i/(numRows-1)%2 == 0 {
			seg[index] += string(v)
		} else {
			seg[numRows-index-1] += string(v)
		}
	}

	var ret string
	for _, x := range seg {
		ret += string(x)
	}

	return ret
}
