/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
package main

import "fmt"

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	lens, lent := len(s), len(t)
	dp := make([][]int, lens+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, lent+1)
	}

	for i := 0; i <= lens; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= lens; i++ {
		for j := 1; j <= lent; j++ {
			dp[i][j] = dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[lens][lent]
}

func main() {
	fmt.Println(numDistinct("ab", "c"))
}

// @lc code=end
