/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum = targetSum - root.Val

	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if hasPathSum(root.Left, targetSum) {
		return true
	}

	if hasPathSum(root.Right, targetSum) {
		return true
	}
	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			q := queue[i]
			if q.Val == targetSum && q.Left == nil && q.Right == nil {
				return true
			}
			if q.Left != nil {
				q.Left.Val += q.Val
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				q.Right.Val += q.Val
				queue = append(queue, q.Right)
			}
		}
		queue = queue[l:]
	}
	return false
}

// @lc code=end
