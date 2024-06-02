package main

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
//单调栈和哨兵
func maximalRectangle(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])
	height := make([]int, col+2)
	maxArea := 0
	for i := 0; i < row; i++ {
		arr := make([]int, 0)
		arr = append(arr, 0)

		for j := 0; j < col+2; j++ {
			if j > 0 && j <= col {
				if matrix[i][j-1] == '1' {
					height[j] += 1
				} else {
					height[j] = 0
				}
			}

			for len(arr) > 0 && height[arr[len(arr)-1]] > height[j] {
				cur := arr[len(arr)-1]
				arr = arr[:len(arr)-1]

				h := height[cur]
				l := j - arr[len(arr)-1] - 1
				maxArea = maxV1(maxArea, h*l)
			}
			arr = append(arr, j)
		}
	}
	return maxArea
}

func maxV1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minV1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
//动态规划
func maximalRectangle1(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])
	height := make([]int, col)
	rightLess := make([]int, col)
	leftLess := make([]int, col)

	for j := 0; j < col; j++ {
		rightLess[j] = col
		leftLess[j] = -1
	}

	maxArea := 0
	for i := 0; i < row; i++ {

		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				height[j] += 1
			} else {
				height[j] = 0
			}
		}

		boundary := -1
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				leftLess[j] = maxV1(leftLess[j], boundary)
			} else {
				leftLess[j] = -1
				boundary = j
			}
		}

		boundary = col
		for j := col - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				rightLess[j] = minV1(rightLess[j], boundary)
			} else {
				rightLess[j] = col
				boundary = j
			}
		}

		for j := 0; j < col; j++ {
			maxArea = maxV1(maxArea, (rightLess[j]-leftLess[j]-1)*height[j])
		}

	}
	return maxArea
}
