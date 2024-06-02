package main

/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	res := make([][]int, n1+1)
	for i, _ := range res {
		res[i] = make([]int, n2+1)
	}
	res[0][0] = 1
	for i := 1; i <= n1; i++ {
		if res[i-1][0] > 0 && s1[i-1] == s3[i-1] {
			res[i][0] = 1
		}
	}
	for i := 1; i <= n2; i++ {
		if res[0][i-1] > 0 && s2[i-1] == s3[i-1] {
			res[0][i] = 1
		}
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if res[i][j-1] > 0 && s3[i+j-1] == s2[j-1] {
				res[i][j] = 1
			}
			if res[i-1][j] > 0 && s3[i+j-1] == s1[i-1] {
				res[i][j] = 1
			}
		}
	}
	return res[n1][n2] == 1
}

// @lc code=end
