package main

import "math"

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return dividend
	}
	if divisor == 1 {
		return dividend
	}
	pre := 1
	if dividend > 0 && divisor < 0 {
		pre = -1
	}
	if dividend < 0 && divisor > 0 {
		pre = -1
	}
	a, b := dividend, divisor
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	count := mul(a, b)
	return pre * count
}

func mul(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	count := 1
	tmp := divisor
	for (tmp << 1) <= dividend {
		count <<= 1
		tmp <<= 1
	}

	return count + mul(dividend-tmp, divisor)
}

// @lc code=end
