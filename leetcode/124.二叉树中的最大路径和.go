/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {

	maxValue := math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		cur := left + right + root.Val
		maxValue = maxV(cur, maxValue)
		ret := maxV(left, right) + root.Val

		return maxV(ret, 0)
	}
	dfs(root)
	return maxValue
}

func maxV(a, b int) int {
	if a > b {
		return a
	}
	return b
}