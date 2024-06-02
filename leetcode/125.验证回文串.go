/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
package main

import "strings"

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	str := strings.ToLower(s)
	left, right := 0, len(str)-1
	for left < right {
		if !(str[left] >= 'a' && str[left] <= 'z' || str[left] >= '0' && str[left] <= '9') {
			left++
			continue
		}
		if !(str[right] >= 'a' && str[right] <= 'z' || str[right] >= '0' && str[right] <= '9') {
			right--
			continue
		}
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end
