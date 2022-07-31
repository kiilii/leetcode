package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 86.62% 的用户
func generate(numRows int) [][]int {
	var ret = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		for ii := range ret[i] {
			if ii == 0 || ii == len(ret[i])-1 {
				ret[i][ii] = 1
			} else {
				ret[i][ii] = ret[i-1][ii-1] + ret[i-1][ii]
			}
		}
	}
	return ret
}
