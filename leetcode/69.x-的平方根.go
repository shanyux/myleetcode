package main

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 1, x
	for left < right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

// @lc code=end
