package main

// 2023.05.29
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 96.45% 的用户
func plusOne(digits []int) []int {
	var length = len(digits)
	digits[length-1] += 1

	for i := length - 1; i > 0; i-- {
		if digits[i]/10 > 0 {
			digits[i] = digits[i] % 10
			digits[i-1] += 1
		}
	}

	if digits[0] == 10 {
		digits = append(digits, 0)
		digits[0] = 0
		for ; length > 0; length-- {
			digits[length] = digits[length-1]
		}
		digits[0] = 1
	}

	return digits
}
