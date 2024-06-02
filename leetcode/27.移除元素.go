package main

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	j := len(nums) - 1
	for i := 0; i <= j; i++ {
		if nums[i] != val {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i--
	}
	return j + 1
}

// @lc code=end
