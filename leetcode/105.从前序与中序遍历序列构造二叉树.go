/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	return buildTreeStruct(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildTreeStruct(preorder []int, prestart, preend int, inorder []int, instart, inend int) *TreeNode {
	if prestart > preend {
		return nil
	}

	root := &TreeNode{Val: preorder[prestart]}
	mid := 0
	for i := instart; i <= inend; i++ {
		if inorder[i] == root.Val {
			mid = i
			break
		}
	}
	l := mid - instart
	root.Left = buildTreeStruct(preorder, prestart+1, prestart+l, inorder, instart, mid-1)
	root.Right = buildTreeStruct(preorder, prestart+1+l, preend, inorder, mid+1, inend)
	return root
}

// @lc code=end
