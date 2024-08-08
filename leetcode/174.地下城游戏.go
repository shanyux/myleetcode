/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
package main

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == m-1 && j == n-1 {
			if dungeon[i][j] >= 0 {
				return 1
			} else {
				return 1 - dungeon[i][j]
			}
		}

		if i >= m || j >= n {
			return math.MaxInt
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		t := dp(i+1, j)
		if t > dp(i, j+1) {
			t = dp(i, j+1)
		}
		memo[i][j] = t - dungeon[i][j]
		if memo[i][j] <= 0 {
			memo[i][j] = 1
		}
		return memo[i][j]
	}
	return dp(0, 0)
}

func calculateMinimumHP1(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n+1)

	}

	for j := 0; j <= n; j++ {
		memo[m][j] = math.MaxInt
	}
	for i := 0; i <= m; i++ {
		memo[i][n] = math.MaxInt
	}
	memo[m-1][n-1] = 1
	if dungeon[m-1][n-1] < 0 {
		memo[m-1][n-1] = 1 - dungeon[m-1][n-1]
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}

			minV := memo[i+1][j]
			if minV > memo[i][j+1] {
				minV = memo[i][j+1]
			}
			t := minV - dungeon[i][j]
			memo[i][j] = t
			if t <= 0 {
				memo[i][j] = 1
			}
		}
	}
	return memo[0][0]

}

// @lc code=end
