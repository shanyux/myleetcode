package main

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	res := build(1, n)
	return res

}

func build(lo, hi int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for i := lo; i <= hi; i++ {
		left := build(lo, i-1)
		right := build(i+1, hi)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}

// @lc code=end
