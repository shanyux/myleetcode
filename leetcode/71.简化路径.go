package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	tmp := make([]string, 0)
	for _, s := range strs {
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			if len(tmp) > 0 {
				tmp = tmp[:len(tmp)-1]
			}
		} else {
			tmp = append(tmp, s)
		}
	}
	res := ""
	for _, s := range tmp {
		res = res + "/" + s
	}
	if res == "" {
		res = "/"
	}
	return res

}

// @lc code=end
