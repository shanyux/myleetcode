package main

/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	a, b := numerator, denominator
	if a%b == 0 {
		return string(a / b)
	}
	isFu := false
	if a*b < 0 {
		isFu = true
	}
	var res string

}

// @lc code=end
