/*
 * @lc app=leetcode.cn id=201 lang=golang
 * @lcpr version=20000
 *
 * [201] 数字范围按位与
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	mask := 1 << 30
	res := 0
	for mask != 0 && (left&mask) == (right&mask) {
		res |= left & mask
		mask >>= 1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 5\n7\n
// @lcpr case=end

// @lcpr case=start
// 0\n0\n
// @lcpr case=end

// @lcpr case=start
// 1\n2147483647\n
// @lcpr case=end

*/

