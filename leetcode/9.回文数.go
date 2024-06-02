package opsappservice

import ()

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	res := 0
	for tmp != 0 {
		last := tmp % 10
		tmp = tmp / 10
		res = res*10 + last
	}
	return x == res
}

// @lc code=end
