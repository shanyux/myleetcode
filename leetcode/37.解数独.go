package main

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, i, j int) bool {
	m, n := 9, 9
	if j >= n {
		return backtrack(board, i+1, 0)
	}
	if i >= m {
		return true
	}

	if board[i][j] != '.' {
		return backtrack(board, i, j+1)
	}

	for c := '1'; c <= '9'; c++ {
		if !isValidSudu(board, byte(c), i, j) {
			continue
		}
		board[i][j] = byte(c)
		if backtrack(board, i, j+1) {
			return true
		}

		board[i][j] = '.'

	}
	return false
}

func isValidSudu(board [][]byte, b byte, r, c int) bool {
	for k := 0; k < 9; k++ {
		if board[r][k] == b {
			return false
		}
		if board[k][c] == b {
			return false
		}
		if board[r/3*3+k/3][c/3*3+k%3] == b {
			return false
		}
	}
	return true
}

// @lc code=end
