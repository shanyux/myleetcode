package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	pre := strs[0]
	for _, str := range strs[1:] {
		temp := ""
		for i, _ := range str {
			if i >= len(pre) {
				break
			}
			if str[i] != pre[i] {
				break
			}
			temp = str[0 : i+1]
			// fmt.Println(temp)
		}
		pre = temp
	}
	return pre
}

func main() {
	fmt.Print(longestCommonPrefix([]string{"cir", "car"}))
}

// @lc code=end
