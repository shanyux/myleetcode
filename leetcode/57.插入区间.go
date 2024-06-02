package main

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	need := true

	res := make([][]int, 0)
	tmp := newInterval
	for i := 0; i < len(intervals); i++ {
		if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else {
			tmp[0] = minNum5(tmp[0], intervals[i][0])
		}

		if newInterval[1] < intervals[i][0] {
			if need {
				res = append(res, tmp)
				need = false
			}
			res = append(res, intervals[i])
		} else {
			tmp[1] = maxNum5(intervals[i][1], tmp[1])
		}

	}
	if need {
		res = append(res, tmp)
	}

	return res
}

func maxNum5(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minNum5(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
