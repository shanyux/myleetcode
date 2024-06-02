package main

import "strings"

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	backtrack1(board, 0, &res)
	return res
}

func backtrack1(board []string, row int, res *[][]string) {
	if row == len(board) {
		*res = append(*res, append([]string{}, board...))
		return
	}
	n := len(board)
	for i := 0; i < n; i++ {
		if !isValid1(board, row, i) {
			continue
		}
		str := []byte(board[row])
		str[i] = 'Q'
		board[row] = string(str)
		backtrack1(board, row+1, res)
		str[i] = '.'
		board[row] = string(str)

	}
}

func isValid1(board []string, row, col int) bool {
	for i := 0; i < row; i++ {
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
