package main

import "strings"

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func isNumber(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	numberSee, eSee, pointSee, numberAfterE := false, false, false, false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numberSee = true
			numberAfterE = true
		} else if s[i] == '.' {
			if eSee || pointSee {
				return false
			}
			pointSee = true
		} else if s[i] == 'e' {
			if eSee || !numberSee {
				return false
			}
			eSee = true
			numberAfterE = false
		} else if s[i] == '+' || s[i] == '-' {
			if i != 0 && s[i-1] != 'e' {
				return false
			}
		} else {
			return false
		}
	}
	return numberSee && numberAfterE
}

// @lc code=end
