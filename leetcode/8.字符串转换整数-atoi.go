package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return 0
	}
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	if i == len(s) {
		return 0
	}
	res := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		if res > math.MaxInt32 {
			break
		}
		i++
	}
	if res > math.MaxInt32 {
		if sign > 0 {
			return math.MaxInt32
		}
		return math.MinInt32
	}
	return sign * res
}

func main() {
	fmt.Print(myAtoi("9223372036854775808"))
}

// @lc code=end
