package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s1 string) string {
	str := strings.TrimSpace(s1)
	tmpArrive := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		if i < len(str)-1 && str[i+1] == ' ' {
			continue
		}
		tmpArrive = append(tmpArrive, str[i])
	}
	pre := 0
	for i := 0; i < len(tmpArrive); i++ {
		if tmpArrive[i] == ' ' {
			reverseWords1(tmpArrive, pre+1, i-1)
		}
		pre = i
	}
	return string(tmpArrive)
}

func reverseWords1(arr []byte, i, j int) {
	for i < j {
		t := arr[i]
		arr[i] = arr[j]
		arr[j] = t
		i++
		j--
	}
}

// @lc code=end
