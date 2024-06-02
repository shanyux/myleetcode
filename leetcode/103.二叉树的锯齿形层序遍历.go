/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	flag := 0
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < l; i++ {
			q := queue[0]
			if flag%2 == 0 {
				tmp = append(tmp, q.Val)
			} else {
				tmp = append([]int{q.Val}, tmp...)
			}
			queue = queue[1:]
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		res = append(res, tmp)
		flag++
	}
	return res
}

// @lc code=end
