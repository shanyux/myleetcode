package main

import "strconv"

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	res := "1"
	for i := 1; i < n; i++ {
		c := res[0]
		tmp := ""
		num := 0
		for j := 0; j < len(res); j++ {
			if c == res[j] {
				num++
				continue
			}
			tmp = tmp + strconv.Itoa(num) + string(c)
			c = res[j]
			num = 1
		}
		res = tmp + strconv.Itoa(num) + string(c)
	}
	return res
}

// @lc code=end
