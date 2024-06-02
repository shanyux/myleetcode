package main

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	return leftBound1(nums, target)
}

func leftBound1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		}
		if target < nums[mid] {
			right = mid - 1
		}
	}

	return left
}

// @lc code=end
