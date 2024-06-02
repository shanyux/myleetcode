/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	wordMap := make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}
	level := 1
	queue := []string{beginWord}

	for len(queue) > 0 {
		l := len(queue)
		for _, str := range queue[:l] {
			if str == endWord {
				return level
			}
			for i := 0; i < len(str); i++ {
				for j := 'a'; j <= 'z'; j++ {
					new := str[:i] + string(j) + str[i+1:]

					if wordMap[new] {
						wordMap[new] = false
						queue = append(queue, new)
					}
				}
			}
		}
		level++
		queue = queue[l:]
	}
	return 0
}

// @lc code=end
