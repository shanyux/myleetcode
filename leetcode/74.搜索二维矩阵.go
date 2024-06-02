package main

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix1(matrix [][]int, target int) bool { //二分
	r, l := len(matrix), len(matrix[0])
	left, right := 0, r*l-1
	for left <= right {
		mid := left + (right-left)/1
		i, j := getIJ(r, l, mid)
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func getIJ(r, l, index int) (int, int) {
	i := index / l
	j := index % l
	return i, j
}

func searchMatrix(matrix [][]int, target int) bool { // BST二叉查找树
	r, l := len(matrix), len(matrix[0])
	x, y := 0, l-1
	for x >= 0 && x < r && y >= 0 && y < l {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}

	return false
}

// @lc code=end
