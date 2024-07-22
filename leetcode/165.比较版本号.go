/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	str1, str2 := strings.Split(version1, "."), strings.Split(version2, ".")
	l1, l2, i, j := len(str1), len(str2), 0, 0
	for i < l1 || j < l2 {
		a, b := 0, 0
		if i < l1 {
			a, _ = strconv.Atoi(str1[i])
			i++
		}
		if j < l2 {
			b, _ = strconv.Atoi(str2[j])
			j++
		}
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
	}
	return 0
}

// @lc code=end
