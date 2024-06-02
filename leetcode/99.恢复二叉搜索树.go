package opsappservice

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
//碰到大小不对的node对<a,b>  前面的取a 后面的取b
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var pre, x, y *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pre != nil {
			if pre.Val > root.Val {
				y = root
				if x == nil {
					x = pre
				}
			}
		}

		pre = root
		dfs(root.Right)
	}
	dfs(root)
	if x != nil && y != nil {
		v := x.Val
		x.Val = y.Val
		y.Val = v
	}

}

// @lc code=end
