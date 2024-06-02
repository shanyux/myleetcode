package main

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	if len(p) > len(s) {
		return nil
	}
	windows, needs := make(map[byte]int, 0), make(map[byte]int, 0)
	for i, _ := range p {
		needs[p[i]]++
	}
	valid := 0
	res := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if needs[c] == windows[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(needs) {
				res = append(res, left)
			}
			d := s[left]
			if _, ok := needs[d]; ok {
				if needs[d] == windows[d] {
					valid--
				}
				windows[d]--
			}
			left++
		}
	}
	return res
}

// @lc code=end
