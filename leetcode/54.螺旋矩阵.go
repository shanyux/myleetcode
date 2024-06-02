package main

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n)

	upper_board := 0
	lower_board := m - 1
	left_board := 0
	right_board := n - 1

	for len(res) < m*n {
		if upper_board <= lower_board {
			for i := left_board; i <= right_board; i++ {
				res = append(res, matrix[upper_board][i])
			}
			upper_board++
		}

		if left_board <= right_board {
			for i := upper_board; i <= lower_board; i++ {
				res = append(res, matrix[i][right_board])
			}
			right_board--
		}

		if upper_board <= lower_board {
			for i := right_board; i >= left_board; i-- {
				res = append(res, matrix[lower_board][i])
			}
			lower_board--
		}

		if left_board <= right_board {
			for i := lower_board; i >= upper_board; i-- {
				res = append(res, matrix[i][left_board])
			}
			left_board++
		}

	}
	return res

}

// @lc code=end
