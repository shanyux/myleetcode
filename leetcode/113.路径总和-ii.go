/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func pathSum1(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	tmp := make([]int, 0)
	var backTrank func(root *TreeNode, tmp []int, targetSum int)
	backTrank = func(root *TreeNode, tmp []int, targetSum int) {
		if root == nil {
			return
		}
		targetSum = targetSum - root.Val
		tmp = append(tmp, root.Val)

		if targetSum == 0 && root.Left == nil && root.Right == nil {
			t := make([]int, len(tmp))
			copy(t, tmp)
			res = append(res, t)
			return
		}

		backTrank(root.Left, tmp, targetSum)

		backTrank(root.Right, tmp, targetSum)

		tmp = tmp[:len(tmp)-1]
	}
	backTrank(root, tmp, targetSum)

	return res
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	tmp := make([]int, 0)
	queue := make([]*TreeNode, 0)
	list := make([][]int, 0)

	queue = append(queue, root)
	tmp = append(tmp, root.Val)
	root.Val = targetSum - root.Val
	list = append(list, tmp)

	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			q := queue[i]
			temp := list[i]
			t := make([]int, len(temp))
			copy(t, temp)
			if q.Val == 0 && q.Left == nil && q.Right == nil {
				res = append(res, t)
				continue
			}

			if q.Left != nil {
				queue = append(queue, q.Left)
				// t = append(t, q.Left.Val)
				list = append(list, append(t, q.Left.Val))
				q.Left.Val = q.Val - q.Left.Val
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
				// t = append(t, q.Right.Val)
				list = append(list, append(t, q.Right.Val))
				q.Right.Val = q.Val - q.Right.Val

			}
		}
		queue = queue[l:]
		list = list[l:]
	}

	return res
}

// @lc code=end

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	tmp := make([]int, 0)
	queue := make([]*TreeNode, 0)
	list := make([][]int, 0)

	queue = append(queue, root)
	tmp = append(tmp, root.Val)
	list = append(list, tmp)

	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			q := queue[i]
			temp := list[i]
			t := make([]int, len(temp))
			copy(t, temp)
			if q.Val == targetSum && q.Left == nil && q.Right == nil {
				res = append(res, t)
				continue
			}

			if q.Left != nil {
				list = append(list, append(t, q.Left.Val))
				q.Left.Val = q.Val + q.Left.Val
				queue = append(queue, q.Left)
				// t = append(t, q.Left.Val)
			}
			if q.Right != nil {

				q.Right.Val = q.Val + q.Right.Val
				queue = append(queue, q.Right)
				list = append(list, append(t, q.Right.Val))

				// t = append(t, q.Right.Val)

			}
		}
		queue = queue[l:]
		list = list[l:]
	}

	return res
}
