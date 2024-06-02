package main

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	//https://leetcode.cn/problems/regular-expression-matching/solutions/296114/shou-hui-tu-jie-wo-tai-nan-liao-by-hyj8/

	//abcda
	//abc*a
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2] //*前面出现0次或 s[i−1] 和 p[j−2] 不匹配
				if match(s, p, i, j-1) {          //s[i−1] 和 p[j−2]匹配 i要去上取一个
					dp[i][j] = dp[i][j] || dp[i-1][j] || dp[i-1][j-2] // 多次 1次
				}
			} else if match(s, p, i, j) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func match(s, p string, i, j int) bool {
	if i == 0 {
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}

// @lc code=end
