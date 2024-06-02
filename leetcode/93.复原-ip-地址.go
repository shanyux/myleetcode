package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	res, tmp := make([]string, 0), make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) && len(tmp) == 4 {
			res = append(res, strings.Join(tmp, "."))
			return
		}
		for i := start; i < len(s); i++ {
			if !isValid3(s[start : i+1]) {
				continue
			}

			if len(tmp) >= 4 {
				break
			}

			tmp = append(tmp, s[start:i+1])
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0)
	return res
}

func isValid3(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if s[0] == '0' {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if v < 0 || v > 255 {
		return false
	}

	return true
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

// @lc code=end
