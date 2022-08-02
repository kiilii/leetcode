package main

// 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了 77.39% 的用户
func getRow(rowIndex int) []int {
	var ret = [][]int{{}, {}}

	for i := 0; i < rowIndex+1; i++ {
		for ii := 0; ii < i+1; ii++ {
			if ii == 0 || ii == i {
				ret[1] = append(ret[1], 1)
			} else {
				ret[1] = append(ret[1], ret[0][ii-1]+ret[0][ii])
			}
		}

		ret[0] = ret[0][:0]
		ret[0] = append(ret[0], ret[1]...)
		ret[1] = ret[1][:0]
	}
	return ret[0]
}

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 39.76% 的用户
func getRow2(rowIndex int) []int {
	var ret = make([][]int, rowIndex+1)

	for i := 0; i < rowIndex+1; i++ {
		ret[i] = make([]int, i+1)
		for ii := range ret[i] {
			if ii == 0 || ii == len(ret[i])-1 {
				ret[i][ii] = 1
			} else {
				ret[i][ii] = ret[i-1][ii-1] + ret[i-1][ii]
			}
		}
	}
	return ret[rowIndex]
}

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 67.69% 的用户
func getRow3(rowIndex int) []int {
	var ret = make([][]int, 2)
	ret[0] = []int{1}

	for i := 0; i < rowIndex+1; i++ {
		ret[1] = make([]int, i+1)
		for ii := range ret[1] {
			if ii == 0 || ii == i {
				ret[1][ii] = 1
			} else {
				ret[1][ii] = ret[0][ii-1] + ret[0][ii]
			}
		}
		ret[0] = ret[1]
	}
	return ret[0]
}
