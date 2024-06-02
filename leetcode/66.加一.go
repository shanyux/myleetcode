package main

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	res := make([]int, 0, l+1)
	res[0] = 1
	res = append(res, digits...)
	return res
}

// @lc code=end
