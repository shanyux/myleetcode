/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(0, len(nums)-1, nums)
}

func dfs(l, r int, nums []int) *TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = dfs(l, mid-1, nums)
	root.Right = dfs(mid+1, r, nums)
	return root
}

// @lc code=end
