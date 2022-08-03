package main

import "math"

// 执行用时：100 ms, 在所有 Go 提交中击败了 71.91% 的用户
// 内存消耗：7.8 MB, 在所有 Go 提交中击败了 67.46% 的用户
func maxProfit(prices []int) int {
	var profit = 0
	var maxNum = math.MinInt32

	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxNum {
			maxNum = prices[i]
		} else {
			profit = max(profit, maxNum-prices[i])
		}
	}

	return profit
}
