package main

/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
// https://leetcode.cn/problems/gray-code/solutions/13637/gray-code-jing-xiang-fan-she-fa-by-jyd/
// 算法很简单 可惜记不住

func grayCode(n int) []int {
	head := 1
	res := []int{0}
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]+head)
		}
		head <<= 1
	}
	return res
}

// @lc code=end
