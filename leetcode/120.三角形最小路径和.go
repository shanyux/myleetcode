/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
package main

var memo [][]*int

func minimumTotal(triangle [][]int) int {
	r := len(triangle)
	memo = make([][]*int, r)
	for i, _ := range memo {
		memo[i] = make([]*int, r)
	}
	return dfs5(triangle, 0, 0)
}

func dfs5(triangle [][]int, i, j int) int {
	if i == len(triangle) {
		return 0
	}
	if memo[i][j] != nil {
		return *memo[i][j]
	}
	a := dfs5(triangle, i+1, j)
	b := dfs5(triangle, i+1, j+1)
	var res int
	if a > b {
		res = b + triangle[i][j]
	} else {
		res = a + triangle[i][j]
	}
	memo[i][j] = &res
	return *memo[i][j]
}

// @lc code=end

func minimumTotal(triangle [][]int) int {
	r := len(triangle)
	dp := make([][]int, r+1)
	for i, _ := range dp {
		dp[i] = make([]int, r+1)
	}

	for i := r - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			t := 0
			if dp[i+1][j] > dp[i+1][j+1] {
				t = dp[i+1][j+1]
			} else {
				t = dp[i+1][j]
			}
			dp[i][j] = t + triangle[i][j]
		}
	}
	return dp[0][0]
}
