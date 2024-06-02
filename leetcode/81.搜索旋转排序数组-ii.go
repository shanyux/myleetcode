package main

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search1(nums []int, target int) bool {
	left, right := 0, len(nums)

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
			continue
		}
		if nums[right] == nums[mid] {
			right--
			continue
		}
		if nums[mid] > nums[left] {
			if nums[mid] > target && nums[left] < target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] > target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// @lc code=end
