/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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

import "math"

// 叶子节点的定义是左孩子和右孩子都为 null 时叫做叶子节点
// 当 root 节点左右孩子都为空时，返回 1
// 当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
// 当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值

// 作者：房建斌学算法
// 链接：https://leetcode.cn/problems/minimum-depth-of-binary-tree/solutions/11486/li-jie-zhe-dao-ti-de-jie-shu-tiao-jian-by-user7208/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	l, r := math.MaxInt, math.MaxInt
	if root.Left != nil {
		l = minDepth(root.Left)
	}
	if root.Right != nil {
		r = minDepth(root.Right)
	}
	if l >= r {
		return r + 1
	}
	return l + 1
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := minDepth(root.Left)

	r := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return l + r + 1
	}
	if l >= r {
		return r + 1
	}
	return l + 1
}

// @lc code=end
