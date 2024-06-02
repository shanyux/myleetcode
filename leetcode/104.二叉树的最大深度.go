/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxV4(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxV4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		res++
		l := len(queue)
		for i := 0; i < l; i++ {
			q := queue[0]
			queue = queue[1:]

			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}

	}

	return res
}

// @lc code=end
