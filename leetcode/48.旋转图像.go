package main

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		k := 0
		j := n - 1
		for k < j {
			matrix[i][k], matrix[i][j] = matrix[i][j], matrix[i][k]
			k++
			j--
		}
	}

}

// @lc code=end
