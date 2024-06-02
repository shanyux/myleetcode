package main

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	mid := 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] == target {
			hi = mid - 1
		}
		if target > nums[mid] {
			lo = mid + 1
		}
		if target < nums[mid] {
			hi = mid - 1
		}
	}
	if lo >= len(nums) {
		return -1
	}
	if nums[lo] != target {
		return -1
	}
	return lo
}

func rightBound(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	mid := 0
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if nums[mid] == target {
			lo = mid + 1
		}
		if target > nums[mid] {
			lo = mid + 1
		}
		if target < nums[mid] {
			hi = mid - 1
		}
	}
	if hi < 0 {
		return -1
	}
	if nums[hi] != target {
		return -1
	}
	return hi

}

// @lc code=end

// @lc code=end
