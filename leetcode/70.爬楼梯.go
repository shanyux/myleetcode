package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	a1, a2, res := 1, 2, 0
	for i := 3; i <= n; i++ {
		res = a1 + a2
		a1 = a2
		a2 = res

	}
	return res
}

// @lc code=end
