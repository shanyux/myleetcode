/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
package main

func isBalanced(root *TreeNode) bool {
	if getDepth(root) == -1 {
		return false
	}
	return true
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := getDepth(root.Left)
	if l == -1 {
		return -1
	}
	r := getDepth(root.Right)
	if r == -1 {
		return -1
	}

	if r-l > 1 {
		return -1
	}
	if l-r > 1 {
		return -1
	}
	return maxNum6(l, r) + 1
}

func maxNum6(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
