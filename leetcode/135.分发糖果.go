package main

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	l := len(ratings)
	if l == 0 {
		return 0
	}
	left := make([]int, l)
	right := make([]int, l)

	for i := 0; i < l; i++ {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 0; i < l; i++ {
		if left[i] > right[i] {
			res += left[i]
		} else {
			res += right[i]
		}
	}
	return res
}

// @lc code=end
