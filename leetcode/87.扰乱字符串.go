package main

/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
//递归
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	if !check(s1, s2) {
		return false
	}

	n := len(s1)
	for k := 1; k < len(s1)-1; k++ {
		a, b := s1[:k], s1[k:]
		c, d := s2[:k], s2[k:]
		if isScramble(a, c) && isScramble(b, d) {
			return true
		}
		e, f := s2[n-k:], s2[:n-k]
		if isScramble(a, e) && isScramble(b, f) {
			return true
		}
	}

	return false
}

//递归+剪枝
func isScramble1(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	if !check(s1, s2) {
		return false
	}
	n := len(s1)
	cache := make([][][]int, n)
	for i, _ := range cache {
		cache[i] = make([][]int, n)
		for j, _ := range cache[i] {
			cache[i][j] = make([]int, n+1)
		}
	}
	var dfs func(i, j, len int) bool
	dfs = func(i, j, len int) bool {
		if cache[i][j][len] > 0 {
			return cache[i][j][len] == 1
		}
		a := s1[i : i+len]
		b := s2[j : j+len]
		if a == b {
			cache[i][j][len] = 1
			return true
		}

		if !check(a, b) {
			cache[i][j][len] = 2
			return false
		}

		for k := 1; k < len; k++ {
			if dfs(i, j, k) && dfs(i+k, j+k, len-k) {
				cache[i][j][len] = 1
				return true
			}
			if dfs(i, j+len-k, k) && dfs(i+k, j, len-k) {
				cache[i][j][len] = 1
				return true
			}
		}
		cache[i][j][len] = 2
		return false
	}
	return dfs(0, 0, n)
}

// @lc code=end

func check(s1, s2 string) bool {
	charCount1, charCount2 := make(map[int]int, 0), make(map[int]int, 0)
	for i, _ := range s1 {
		charCount1[int(s1[i]-'a')]++
	}
	for i, _ := range s2 {
		charCount2[int(s2[i]-'a')]++
	}

	for i := 0; i < 26; i++ {
		if charCount1[i] != charCount2[i] {
			return false
		}
	}
	return true

}

//动态规划 不好懂
func isScramble2(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}

	n := len(s1)
	cache := make([][][]int, n)
	for i, _ := range cache {
		cache[i] = make([][]int, n)
		for j, _ := range cache[i] {
			cache[i][j] = make([]int, n+1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s1[i] == s2[j] {
				cache[i][j][1] = 1
			}
		}
	}

	for len := 2; len <= n; len++ {
		for i := 0; i <= n-len; i++ { //len个字符一起判断  所以要n-len
			for j := 0; j <= n-len; j++ {
				for k := 1; k < len; k++ {
					if cache[i][j][k] == 1 && cache[i+k][j+k][len-k] == 1 {
						cache[i][j][len] = 1
					}
					if cache[i][len-k+j][k] == 1 && cache[i+k][j][len-k] == 1 {
						cache[i][j][len] = 1
					}
				}
			}
		}
	}

	return cache[0][0][n] == 1
}
