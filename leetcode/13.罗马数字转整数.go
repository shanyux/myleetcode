package main

import ()

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	res := 0
	CharValueMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	bytes := []byte(s)
	for i, c := range bytes {
		value := CharValueMap[c]
		if i < len(s)-1 && value < CharValueMap[bytes[i+1]] {
			res = res - value
		} else {
			res = res + value
		}
	}
	return res
}

// @lc code=end
