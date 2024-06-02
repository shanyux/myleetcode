package main

import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if !dfs(root.Left) {
			return false
		}

		if root.Val < pre {
			return false
		}
		pre = root.Val
		return dfs(root.Right)
	}
	return dfs(root)
}

// @lc code=end
