/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
package main

func getRow(rowIndex int) []int {
	res := make([]int, 0)
	res = append(res, 1)
	pre := 1
	for i := 1; i <= rowIndex; i++ {
		for j := 1; j < i; j++ {
			t := res[j]
			res[j] = pre + res[j]
			pre = t
		}
		res = append(res, 1)
	}
	return res
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		res[0] = 1
		res[i] = 1
		for j := i - 1; j >= 1; j-- {
			res[j] += res[j-1]
		}

	}
	return res
}

// @lc code=end
