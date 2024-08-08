/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for x >= sx {
			sx = sx * 10
		}

		for y >= sy {
			sy = sy * 10
		}
		if y*sx+x > x*sy+y {
			return false
		}
		return true
	})

	if nums[0] == 0 {
		return "0"
	}
	res := ""
	for _, n := range nums {
		res += strconv.Itoa(n)
	}
	return res
}

// @lc code=end
