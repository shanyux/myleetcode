package main

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	queue := make([]int, 0)
	queue = append(queue, -1)
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			queue = append(queue, i)
		} else {
			queue = queue[:len(queue)-1]
			if len(queue) == 0 {
				queue = append(queue, i)
				continue
			}
			maxLen = maxNum(maxLen, i-queue[len(queue)-1])
		}
	}
	return maxLen
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses1(s string) int {
	if len(s) < 2 {
		return 0
	}
	maxLen := 0

	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-1 > 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-1 > 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxLen = maxNum(maxLen, dp[i])
	}
	return maxLen
}

// @lc code=end
