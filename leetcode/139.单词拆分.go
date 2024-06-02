package main

import "strings"

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	wordMap := map[int]bool{}
	return canBreak(s, wordDict, 0, wordMap)
}

func canBreak(s string, wordDict []string, index int, memo map[int]bool) bool {
	if len(s) == 0 {
		return true
	}
	if e, ok := memo[index]; ok {
		return e
	}
	for i := 0; i < len(wordDict); i++ {
		if !strings.HasPrefix(s, wordDict[i]) {
			continue
		}

		if canBreak(s[len(wordDict[i]):], wordDict, index+len(wordDict[i]), memo) {
			memo[index] = true
			return true
		}
		memo[index] = false
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	wordMap := make(map[string]bool)

	for _, s := range wordDict {
		wordMap[s] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for j := i - 1; i >= 0; j-- {
			if wordMap[s[j:i]] && dp[j] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
