/*
 * @lc app=leetcode.cn id=204 lang=golang
 * @lcpr version=20001
 *
 * [204] 计数质数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func countPrimes(n int) int {
	flags := make([]bool, n)
	for i := 0; i < n; i++ {
		flags[i] = true
	}
	for i := 2; i*i < n; i++ {
		for j := i * i; j < n; j += i {
			flags[j] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if flags[i] {
			count++
		}
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
