/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
package main

func minCut(s string) int {
	if s == "" {
		return 0
	}

	maxlen := len(s)
	dp := make([][]int, maxlen)
	for i, _ := range dp {
		dp[i] = make([]int, maxlen)
	}
	for j := 0; j < maxlen; j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = 1
			} else if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = 1
				} else if dp[i+1][j-1] > 0 {
					dp[i][j] = 1
				}

			}
		}
	}
	res := make([]int, maxlen)
	for r := 0; r < maxlen; r++ {
		if dp[0][r] > 0 {
			res[r] = 0
		} else {
			res[r] = r
			for l := 1; l <= r; l++ {
				if dp[l][r] > 0 {
					res[r] = minNum1(res[r], res[l-1]+1)
				}
			}

		}
	}
	return res[maxlen-1]
}

func minNum1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
