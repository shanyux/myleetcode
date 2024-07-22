/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
package main

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// @lc code=end
