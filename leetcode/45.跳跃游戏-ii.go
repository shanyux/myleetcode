package opsappservice

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {

	end := 0
	maxPos := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = maxNum3(i+nums[i], maxPos)
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}

func jump1(nums []int) int {

	end := 1
	maxPos := 0
	step := 0
	start := 0
	for end < len(nums) {
		for i := start; i < end; i++ {
			maxPos = maxNum3(i+nums[i], maxPos)
		}
		start = end
		end = maxPos + 1
		step++
	}
	return step
}

func maxNum3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
