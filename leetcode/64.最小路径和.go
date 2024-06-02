package main

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		res[i][0] = 300
	}
	for i := 0; i <= n; i++ {
		res[0][i] = 300
	}
	res[1][1] = grid[0][0]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			res[i][j] = minNum6(res[i-1][j], res[i][j-1]) + grid[i-1][j-1]
		}
	}
	return res[m][n]

}

func minNum6(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	res[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		res[i][0] = res[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		res[0][i] = res[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = minNum6(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}
	return res[m-1][n-1]

}
