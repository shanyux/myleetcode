package main

import "math"

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == math.MinInt32 {
		return myPow(1/x, -(n+1)) / x
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	if n%2 == 0 {
		v := myPow(x, n/2)
		return v * v
	}
	return myPow(x, n-1) * x
}

// @lc code=end
