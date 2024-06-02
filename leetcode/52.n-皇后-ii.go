package main

import "strings"

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	res := 0

	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	backtrack2(board, 0, &res)
	return res
}

func backtrack2(board []string, row int, res *int) {
	n := len(board)
	if row == n {
		*res++
		return
	}

	for i := 0; i < n; i++ {
		if !isValid2(board, row, i) {
			continue
		}
		tmp := []byte(board[row])
		tmp[i] = 'Q'
		board[row] = string(tmp)
		backtrack2(board, row+1, res)
		tmp[i] = '.'
		board[row] = string(tmp)
	}
}

func isValid2(board []string, row, col int) bool {
	for i := 0; i < row; i = i + 1 {
		if board[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// @lc code=end
