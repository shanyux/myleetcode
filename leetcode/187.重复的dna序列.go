/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
package main

func findRepeatedDnaSequences(s string) []string {
	set := make(map[string]interface{})
	tmp := make(map[string]interface{})
	for i := 0; i <= len(s)-10; i++ {
		t := s[i : i+10]
		if _, ok := tmp[t]; ok {
			set[t] = struct{}{}
		} else {
			tmp[t] = struct{}{}
		}
	}
	res := make([]string, 0)
	for k, _ := range set {
		res = append(res, k)
	}
	return res
}

// @lc code=end
