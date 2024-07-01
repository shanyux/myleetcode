package main

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOrder(root.Left)
		preOrder(root.Right)
		res = append(res, root.Val)

	}
	preOrder(root)
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverseList(res)
	return res
}

func reverseList(arr []int) {
	if len(arr) < 2 {
		return
	}
	i, j := 0, len(arr)-1
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
}

// @lc code=end
