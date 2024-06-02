package main

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for a, b := j, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}

}

// @lc code=end
