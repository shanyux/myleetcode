package main

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if dfs1(board, word, i, j, 0) {
				return true
			}
		}

	}
	return false
}

func dfs1(board [][]byte, word string, i, j, p int) bool {
	m, n := len(board), len(board[0])
	if p == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= m || j >= n {
		return false
	}

	if board[i][j] != word[p] {
		return false
	}
	c := board[i][j]
	board[i][j] = '0'
	found := dfs1(board, word, i+1, j, p+1) ||
		dfs1(board, word, i-1, j, p+1) ||
		dfs1(board, word, i, j+1, p+1) ||
		dfs1(board, word, i, j-1, p+1)
	board[i][j] = c
	return found

}

// @lc code=end
