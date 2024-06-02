package main

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rlen := len(matrix)
	clen := len(matrix[0])
	row0, clo0 := false, false

	for i := 0; i < rlen; i++ {
		if matrix[i][0] == 0 {
			row0 = true
			break
		}
	}
	for i := 0; i < clen; i++ {
		if matrix[0][i] == 0 {
			clo0 = true
			break

		}
	}
	for i := 1; i < rlen; i++ {
		for j := 1; j < clen; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < rlen; i++ {
		for j := 1; j < clen; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for i := 0; i < rlen; i++ {
			matrix[i][0] = 0
		}
	}
	if clo0 {
		for i := 0; i < clen; i++ {
			matrix[0][i] = 0
		}
	}

}

// @lc code=end
