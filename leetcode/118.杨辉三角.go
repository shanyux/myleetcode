/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, 0)
	pre := []int{1}
	res = append(res, pre)
	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < len(tmp)-1; j++ {
			tmp[j] = pre[j-1] + pre[j]
		}
		res = append(res, tmp)
		pre = tmp
	}
	return res
}

// @lc code=end
