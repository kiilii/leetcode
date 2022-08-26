package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 9.64% 的用户
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var m = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"}}

	var x = make([][]string, 0)

	for _, n := range digits {
		if ss, has := m[string(n)]; has {
			x = append(x, ss)
		}
	}
	return combinations(x)
}

func combinations(x [][]string) []string {
	if len(x) == 1 {
		return x[0]
	}
	ss := combinations(x[1:])
	var ret = make([]string, 0)

	for _, v := range x[0] {
		for _, vv := range ss {
			ret = append(ret, v+vv)
		}
	}
	return ret
}
