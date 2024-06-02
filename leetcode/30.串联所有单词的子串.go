package main

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	l := len(words[0])
	alllen := l * len(words)
	res := make([]int, 0)
	for i := 0; i < len(s)-alllen+1; i++ {
		wordsMap := make(map[string]int)
		value := 0
		for _, str := range words {
			wordsMap[str]++
			value++
		}
		right := i
		for i := 0; i < len(words); i++ {
			tmp := s[right : right+l]
			if wordsMap[tmp] > 0 {
				wordsMap[tmp]--
				value--
			} else {
				break
			}
			right += l
		}
		if value == 0 {
			res = append(res, i)
		}
	}
	return res

}

// @lc code=end
