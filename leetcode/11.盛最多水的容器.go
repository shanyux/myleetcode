package main

import ()

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	maxv := 0
	left, right := 0, len(height)-1
	for left < right {
		h := minVlaue(height[left], height[right])
		newArea := h * (right - left)
		maxv = maxVlaue(maxv, newArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxv

}

func maxVlaue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minVlaue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
