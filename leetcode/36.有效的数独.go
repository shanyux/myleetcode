package main

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {

	row := make([]int, 10)
	col := make([]int, 10)
	area := make([]int, 10)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			u := board[i][j] - '0'
			idx := j/3 + (i/3)*3
			if (row[i]>>u)&1 == 1 || (col[j]>>u)&1 == 1 || (area[idx]>>u)&1 == 1 {
				return false
			}
			row[i] |= 1 << u
			col[j] |= 1 << u
			area[idx] |= 1 << u
		}
	}
	return true
}

// @lc code=end
