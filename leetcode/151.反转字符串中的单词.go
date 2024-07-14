package main

import (
	"fmt"
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
	fmt.Println(string(str))
	tmpArrive := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		if i < len(str)-1 && str[i+1] == ' ' && str[i] == ' ' {
			continue
		}
		tmpArrive = append(tmpArrive, str[i])
	}
	pre := 0
	for i := 0; i < len(tmpArrive); i++ {
		if tmpArrive[i] != ' ' {
			continue
		}
		reverseWords1(tmpArrive, pre, i-1)
		pre = i + 1
	}
	reverseWords1(tmpArrive, pre, len(tmpArrive)-1)
	return string(tmpArrive)
}

func reverseWords1(arr []byte, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

// @lc code=end
