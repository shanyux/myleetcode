package main

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	farest := 0
	for i := 0; i < len(nums); i++ {
		if farest < i {
			return false
		}
		farest = maxNum3(farest, i+nums[i])

	}
	return true
}

func maxNum3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
