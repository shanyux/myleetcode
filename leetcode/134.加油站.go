package opsappservice

import "math"

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return 0
	}
	l := len(gas)
	minV := math.MaxInt
	index := 0
	space := 0

	for i := 0; i < len(gas); i++ {
		space += gas[i] - cost[i]
		if space < minV {
			minV = space
			index = i
		}

	}
	if minV > 0 {
		return 0
	}
	if space < 0 {
		return -1
	}
	return (index + 1) % l
}

// @lc code=end
