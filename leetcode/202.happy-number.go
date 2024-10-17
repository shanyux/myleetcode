/*
 * @lc app=leetcode.cn id=202 lang=golang
 * @lcpr version=20000
 *
 * [202] 快乐数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
package main

func isHappy(n int) bool {
	n1, n2 := n, n
	for {
		n1 = calculate(n1)
		n2 = calculate(n2)
		n2 = calculate(n2)
		if n1 == n2 {
			break
		}
	}
	return n1 == 1
}

func calculate(n int) int {
	num := 0
	for n > 0 {
		v := n % 10
		num += v * v
		n = n / 10
	}
	return num
}

// @lc code=end

/*
// @lcpr case=start
// 19\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
