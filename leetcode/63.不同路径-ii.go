package main

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	res := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int, n+1)
	}
	if obstacleGrid[0][0] == 0 {
		res[1][1] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m][n]

}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	res := make([]int, n+1)

	if obstacleGrid[0][0] == 0 {
		res[1] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				res[j] = 0
				continue
			}
			res[j] = res[j] + res[j-1]
		}
	}
	return res[n]

}

// @lc code=end

// @lc code=end
