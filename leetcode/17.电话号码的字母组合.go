package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}
	res := make([]string, 0)
	path := make([]byte, n)

	dfs(n, 0, digits, &res, &path)
	return res
}

func dfs(maxLen, num int, digits string, res *[]string, path *[]byte) {
	if num >= maxLen {
		*res = append(*res, string(*path))
		return
	}
	numStrings := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	for _, c := range numStrings[digits[num]-'0'] {
		(*path)[num] = byte(c)
		dfs(maxLen, num+1, digits, res, path)
	}

}

// @lc code=end
