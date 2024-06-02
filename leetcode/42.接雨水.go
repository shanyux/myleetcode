package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	lmax, rmax := make([]int, len(height)), make([]int, len(height))
	lmax[0] = 0
	rmax[len(height)-1] = 0
	for i := 1; i < len(height)-1; i++ {
		lmax[i] = maxNum1(lmax[i-1], height[i-1])
	}

	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = maxNum1(rmax[i+1], height[i+1])
	}

	res := 0
	for i := 1; i < len(height)-1; i++ {
		min := minNum1(lmax[i], rmax[i])
		if min > height[i] {
			res += min - height[i]
		}
	}
	return res
}

func maxNum1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minNum1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
