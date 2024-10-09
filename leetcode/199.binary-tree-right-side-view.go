/*
 * @lc app=leetcode.cn id=199 lang=golang
 * @lcpr version=20000
 *
 * [199] 二叉树的右视图
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package myleetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	vals := make([]int, 0)
	tmp := make([]*TreeNode, 0)
	tmp = append(tmp, root)
	for len(tmp) > 0 {
		t := tmp[0]
		l := len(tmp)
		for _, r := range tmp[:l] {
			if r.Left != nil {
				tmp = append(tmp, r.Left)
			}
			if r.Right != nil {
				tmp = append(tmp, r.Right)
			}
			t = r
		}
		tmp = tmp[l:]
		vals = append(vals, t.Val)
	}
	return vals
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
