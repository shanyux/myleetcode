package main

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start

func wordBreak(s string, wordDict []string) []string {
	if len(s) == 0 {
		return nil
	}
	wordMap := make(map[string]bool)
	res := make([]string, 0)
	for _, s := range wordDict {
		wordMap[s] = true
	}

	var dfs func(subs string, tmp string)
	dfs = func(subs string, tmp string) {
		if len(subs) == 0 {
			s1 := tmp
			res = append(res, s1)
			return
		}

		for i := 1; i <= len(subs); i++ {
			t := subs[:i]
			if wordMap[t] {
				if tmp != "" {
					dfs(subs[i:], tmp+" "+t)
				} else {
					dfs(subs[i:], t)
				}
			}
		}
	}
	dfs(s, "")
	return res
}
