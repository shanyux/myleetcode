/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
package main

func majorityElement(nums []int) int {
	cand, count := nums[0], 1
	for i, n := range nums {
		if i == 0 {
			continue
		}
		if n == cand {
			count++
			continue
		}
		count--
		if count == 0 {
			cand = n
			count = 1
		}
	}
	return cand
}

// @lc code=end
