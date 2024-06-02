package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	windows := make(map[byte]int, 0)
	maxlen := 0
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++
		if windows[c] == 1 {
			maxlen = maxLen(maxlen, right-left)
			continue
		}
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
	}
	return maxlen

}

func maxLen(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Print(lengthOfLongestSubstring("bbbbb"))
}

// @lc code=end
