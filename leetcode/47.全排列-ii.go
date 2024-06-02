package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	isUse := make([]bool, len(nums))
	var trackRes func(tmp []int)

	trackRes = func(tmp []int) {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}
		// fmt.Println(tmp)
		preUse := -99999
		for i := 0; i < len(nums); i++ {
			if isUse[i] {
				continue
			}
			// if i > 0 && nums[i] == nums[i-1] && !isUse[i-1] {
			if nums[i] == preUse {
				continue
			}
			preUse = nums[i]
			_ = preUse
			isUse[i] = true
			trackRes(append(tmp, nums[i]))
			isUse[i] = false
		}

	}
	tmp := make([]int, 0)
	trackRes(tmp)
	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1}))
}

// @lc code=end
