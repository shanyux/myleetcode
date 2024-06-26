package main

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil

	// p := root
	// for p.Right != nil {
	// 	p = p.Right
	// }
	// p.Right = right
	if left == nil { //无需处理left
		return
	}

	root.Right = left
	p := left
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

// @lc code=end
