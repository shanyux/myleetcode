package main

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	var dfs func(l, r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}

		if l == nil || r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}
		return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
	}

	return dfs(root.Left, root.Right)
}

func isSymmetric1(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, root)
	for len(queue) > 0 {
		r1 := queue[0]
		r2 := queue[1]
		queue = queue[2:]

		if r1 == nil && r2 == nil {
			continue
		}

		if r1 == nil || r2 == nil {
			return false
		}

		if r1.Val != r2.Val {
			return false
		}
		queue = append(queue, r1.Left)
		queue = append(queue, r2.Right)
		queue = append(queue, r1.Right)
		queue = append(queue, r2.Left)

	}

	return true
}

// @lc code=end
