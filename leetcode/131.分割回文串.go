/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
package main

func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	maxlen := len(s)
	memo := make([][]int, maxlen+1)
	for i, _ := range memo {
		memo[i] = make([]int, maxlen+1)
	}
	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index >= maxlen {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := index; i < maxlen; i++ {
			new := s[index : i+1]
			if memo[index][i] == 2 {
				continue
			}
			if !checkPalindrome(new) {
				memo[index][i] = 2
				continue
			}
			memo[index][i] = 1

			path = append(path, new)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, path)
	return res

}

func checkPalindrome(str string) bool {
	if str == "" {
		return false
	}
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true

}

// @lc code=end
func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	maxlen := len(s)
	dp := make([][]int, maxlen)
	for i, _ := range dp {
		dp[i] = make([]int, maxlen)
	}
	for j := 0; j < maxlen; j++ {
		for i := 0; i <= j; i++ {
			if j == i {
				dp[i][j] = 1
			} else if j-i == 1 && s[j] == s[i] {
				dp[i][j] = 1
			} else if j-i > 1 && s[j] == s[i] && dp[i+1][j-1] > 0 {
				dp[i][j] = 1
			}
		}
	}

	var dfs func(index int, path []string)
	dfs = func(index int, path []string) {
		if index >= maxlen {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := index; i < maxlen; i++ {
			new := s[index : i+1]
			if dp[index][i] < 1 {
				continue
			}

			path = append(path, new)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, path)
	return res

}
