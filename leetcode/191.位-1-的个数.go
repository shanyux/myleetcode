/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
package main

func hammingWeight(n int) int {
	var res int
	for n != 0 {
		if n&1 > 0 {
			res++
		}
		n >>= 1
	}
	return res
}

// @lc code=end
