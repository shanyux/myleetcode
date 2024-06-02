package main

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	left, right := 0, 0
	windows, needs := make(map[byte]int, 0), make(map[byte]int, 0)

	for i, _ := range s1 {
		needs[s1[i]]++
	}

	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if windows[c] == needs[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(needs) {
				return true
			}
			d := s2[left]
			if _, ok := needs[d]; ok {
				if windows[d] == needs[d] {
					valid--
				}
				windows[d]--
			}

			left++
		}

	}
	return false

}

// @lc code=end
