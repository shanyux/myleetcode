package main

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	end := len(s) - 1
	for end >= 0 {
		if s[end] != ' ' {
			break
		}
		end--
	}
	if end < 0 {
		return 0
	}
	start := end

	for start >= 0 {
		if s[start] == ' ' {
			break
		}
		start--
	}

	return end - start
}

// @lc code=end
