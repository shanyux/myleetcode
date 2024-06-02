package main

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1]
		if cur[0] <= last[1] {
			cur[1] = maxNum4(last[1], cur[1])
			continue
		}
		res = append(res, cur)
	}
	return res
}

func maxNum4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
