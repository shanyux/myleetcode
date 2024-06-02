/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minValue := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		minValue = minNum7(minValue, prices[i])
		max = maxNum7(max, prices[i]-minValue)
	}
	return max
}

func maxNum7(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minNum7(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minValue := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		minValue = minNum7(minValue, prices[i])
		max = maxNum7(max, prices[i]-minValue)
	}
	return max
}


// @lc code=end
