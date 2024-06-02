package main

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	num := 1
	upper, left := 0, 0
	low, right := n-1, n-1
	for num <= n*n {
		if upper <= low {
			for i := left; i <= right; i++ {
				res[upper][i] = num
				num++
			}
			upper++
		}
		if left <= right {
			for i := upper; i <= low; i++ {
				res[i][right] = num
				num++
			}
			right--
		}
		if upper <= low {
			for i := right; i >= left; i-- {
				res[low][i] = num
				num++
			}
			low--
		}

		if left <= right {
			for i := low; i >= upper; i-- {
				res[i][left] = num
				num++
			}
			left++
		}
	}
	return res
}

// @lc code=end
