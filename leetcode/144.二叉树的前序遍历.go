package opsappservice

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)

	}
	preOrder(root)
	return res
}

func preorderTraversal1(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			res = append(res, node.Val)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// @lc code=end
