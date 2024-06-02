package main

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0, k)
	var f func(start int, tmp []int)
	f = func(start int, tmp []int) {
		if len(tmp) == k {
			t := make([]int, 0)
			copy(t, tmp)
			res = append(res, t)
			return
		}
		for i := start; i <= n; i++ {
			tmp = append(tmp, i)
			f(i, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	f(1, tmp)
	return res
}

// @lc code=end
