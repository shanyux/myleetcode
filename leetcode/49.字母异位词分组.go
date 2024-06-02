package main

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string][]string, 0)
	for i := 0; i < len(strs); i++ {
		key := encode(strs[i])
		resMap[key] = append(resMap[key], strs[i])
	}
	res := make([][]string, 0)
	for _, value := range resMap {
		res = append(res, value)
	}
	return res
}

func encode(s string) string {
	count := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		d := s[i] - 'a'
		count[d]++
	}
	return string(count)
}

// @lc code=end
