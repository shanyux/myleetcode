/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/solutions/38498/tan-xin-suan-fa-by-liweiwei1419-2/
func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		t := prices[i] - prices[i-1]
		if t > 0 {
			res += t
		}
	}
	return res

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	preBalance, preStock := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		b := maxNum8(preBalance, preStock+prices[i])
		s := maxNum8(preStock, preBalance-prices[i])
		preBalance = b
		preStock = s
	}
	return preBalance

}
func maxNum8(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

