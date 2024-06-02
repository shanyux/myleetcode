package main

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			tmp = append(tmp, nums[i])
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0)
	return res
}

// @lc code=end
