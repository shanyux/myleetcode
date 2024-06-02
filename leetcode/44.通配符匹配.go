package main

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
//https://leetcode.cn/problems/wildcard-matching/solutions/80942/dong-tai-gui-hua-dai-zhu-shi-by-tangweiqun/
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j] // 0字符 1+字符
				continue
			}
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]

			}

		}
	}
	return dp[len(s)][len(p)]
}

// @lc code=end
