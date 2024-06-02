package main

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	arr := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		arr[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		arr[i][0] = arr[i-1][0] + 1
	}
	for i := 1; i <= n; i++ {
		arr[0][i] = arr[0][i-1] + 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				arr[i][j] = arr[i-1][j-1]
			} else {
				arr[i][j] = minNum3(minNum3(arr[i-1][j-1], arr[i][j-1]), arr[i-1][j]) + 1
			}
		}
	}
	return arr[m][n]
}

func minNum3(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
