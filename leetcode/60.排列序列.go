package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	res := ""
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	visit := make([]bool, n+1)

	var backtrack3 func(cur, k int)
	backtrack3 = func(cur, k int) { //cur是下标
		if cur == n {
			return
		}
		num := factorial[n-cur-1]
		for i := 1; i <= n; i++ {
			if visit[i] {
				continue
			}

			if k > num {
				k = k - num
				continue
			}
			res = res + strconv.Itoa(i)
			visit[i] = true
			backtrack3(cur+1, k)
		}
	}
	backtrack3(0, k)
	return string(res)
}

// @lc code=end

func getPermutation1(n int, k int) string {
	res := ""
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	visit := make([]bool, n+1)
	k--
	for i := 1; i <= n; i++ {
		num := factorial[n-i]
		index := (k / num) + 1
		fmt.Println(k, num, index)

		for j := 1; j <= n; j++ {
			if visit[j] {
				continue
			}
			index--
			if index == 0 {
				visit[j] = true
				res = res + strconv.Itoa(j)
				break
			}
		}

		k = k % num
	}
	return string(res)
}

func main() {
	fmt.Println(getPermutation1(3, 2))
}
