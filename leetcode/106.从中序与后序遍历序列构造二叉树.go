/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	return buildTree1(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildTree1(inorder []int, inpre, inend int, postorder []int, postpre, postend int) *TreeNode {
	if postpre > postend || inpre > inend {
		return nil
	}
	val := postorder[postend]

	root := &TreeNode{
		Val: val,
	}
	mid := 0
	for mid = inpre; mid <= inend; mid++ {
		if inorder[mid] == val {
			break
		}
	}
	root.Left = buildTree1(inorder, inpre, mid-1, postorder, postpre, postpre+(mid-inpre)-1)
	root.Right = buildTree1(inorder, mid+1, inend, postorder, postpre+(mid-inpre), postend-1)
	return root
}

// @lc code=end
