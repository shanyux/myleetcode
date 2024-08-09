/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
package main

func convertToTitle(columnNumber int) string {
	res := make([]byte, 0)
	a := columnNumber
	for a > 0 {
		a--
		res = append(res, byte('A'+a%26))
		a /= 26
	}

	return reverseString(res)
}

func reverseString(str []byte) string {
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}

// @lc code=end
