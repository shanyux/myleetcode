package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	left, right := 0, 0
	start := 0
	length := math.MaxInt
	windows, needs := make(map[byte]int), make(map[byte]int)
	for i, _ := range t {
		needs[t[i]]++
	}

	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if windows[c] == needs[c] {
				valid++
			}
		}

		for valid == len(needs) {
			if right-left < length {
				length = right - left
				start = left
			}
			d := s[left]
			left++
			if _, ok := windows[d]; ok {
				if windows[d] == needs[d] {
					valid--
				}
				windows[d]--

			}
		}

	}

	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]

}

// @lc code=end
