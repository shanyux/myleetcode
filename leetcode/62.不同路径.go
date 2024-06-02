package main

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		res[i][0] = 1
	}
	for i := 0; i < n; i++ {
		res[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]

}

func uniquePaths1(m int, n int) int {
	res := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int, n+1)
	}
	res[1][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m][n]

}

// @lc code=end
