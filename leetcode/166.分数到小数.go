package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	a, b := numerator, denominator
	if a%b == 0 {
		return strconv.Itoa(a / b)
	}
	res := ""
	if a*b < 0 {
		res += "-"
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res += strconv.Itoa(a / b)
	hasV := make(map[int]int)
	res += "."
	a = a % b
	for a != 0 {
		hasV[a] = len(res)
		a *= 10
		res += strconv.Itoa(a / b)
		a = a % b
		if index, ok := hasV[a]; ok {
			return fmt.Sprintf("%s(%s)", res[:index], res[index:])
		}
	}
	return res
}

// @lc code=end
