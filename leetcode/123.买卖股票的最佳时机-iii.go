/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	res := make([][]int, n)
	for i, _ := range prices {
		res[i] = make([]int, 4)
	}
	//1买，1卖，2买，2卖
	res[0] = []int{-prices[0], 0, -prices[0], 0}
	for i := 1; i < n; i++ {
		res[i][0] = maxV4(res[i-1][0], -prices[i])
		res[i][1] = maxV4(res[i-1][1], res[i-1][0]+prices[i])
		res[i][2] = maxV4(res[i-1][2], res[i-1][1]-prices[i])
		res[i][3] = maxV4(res[i-1][3], res[i-1][2]+prices[i])
	}
	return res[n-1][4]
}

func maxV4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
