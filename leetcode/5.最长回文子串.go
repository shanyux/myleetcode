package main

import "fmt"

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	maxString := ""
	if len(s) == 1 {
		return s
	}
	for i, _ := range s {
		if i == 0 {
			continue
		}
		sub1 := getSubString(s, i-1, i)
		sub2 := getSubString(s, i, i)
		maxString = getMaxString(maxString, sub1, sub2)
	}
	return maxString
}

func getMaxString(a, b, c string) string {
	m := ""
	if len(m) < len(a) {
		m = a
	}
	if len(m) < len(b) {
		m = b
	}
	if len(m) < len(c) {
		m = c
	}
	return m
}

func getSubString(s string, left, right int) string {
	substring := ""
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			substring = s[left : right+1]
			left--
			right++
			continue
		}
		break
	}
	return substring
}

func main() {
	fmt.Print(longestPalindrome("a"))
}

// @lc code=end
