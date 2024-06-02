package main

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	res := make([]int, n+1)
	res[0] = 1
	if s[0] == '0' {
		res[1] = 0
	} else {
		res[1] = 1
	}
	for i := 2; i <= n; i++ {
		a := s[i-1]
		b := s[i-2]

		if a > '0' && a <= '9' {
			res[i] += res[i-1]
		}

		if b == '1' || (b == '2' && a <= '6') {
			res[i] += res[i-2]
		}
	}
	return res[n]
}

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	a, b, c := 1, 0, 0
	if s[0] != '0' {
		b = 1
		c = 1
	}
	for i := 2; i <= n; i++ {
		c = 0
		d := s[i-1]
		e := s[i-2]

		if d > '0' && d <= '9' {
			c += b
		}

		if e == '1' || (e == '2' && d <= '6') {
			c += a
		}
		a = b
		b = c

	}
	return c
}

// @lc code=end
