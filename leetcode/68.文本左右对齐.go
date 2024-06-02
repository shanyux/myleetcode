package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start
func blank(n int) string {
	return strings.Repeat(" ", n)
}

func fullJustify(words []string, maxWidth int) []string {
	right, n := 0, len(words)
	res := make([]string, 0)
	for {
		sumLen := 0
		left := right
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		if right == n { //最后一行
			tmp := strings.Join(words[left:], " ")
			res = append(res, tmp+blank(maxWidth-len(tmp)))
			return res
		}

		num := right - left
		if num == 1 {
			res = append(res, words[left]+blank(maxWidth-len(words[left])))
			continue
		}

		avgSpace := (maxWidth - sumLen) / (num - 1)
		balanceSpace := (maxWidth - sumLen) % (num - 1)

		leftStr := strings.Join(words[left:left+balanceSpace+1], blank(avgSpace+1))
		rightStr := strings.Join(words[left+balanceSpace+1:right], blank(avgSpace))

		res = append(res, leftStr+blank(avgSpace)+rightStr)
	}
	return res
}

// @lc code=end
