/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

func sortedListToBST(head *ListNode) *TreeNode {
	arr := make([]int, 0)
	h := head
	for h != nil {
		arr = append(arr, h.Val)
		h = h.Next
	}
	return dfs1(arr, 0, len(arr)-1)
}

func dfs1(arr []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	root := &TreeNode{Val: arr[mid]}
	root.Left = dfs1(arr, l, mid-1)
	root.Right = dfs1(arr, mid+1, r)
	return root
}

// @lc code=end

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	var pre *ListNode
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{Val: slow.Val}

	if pre != nil {
		pre.Next = nil
		root.Left = sortedListToBST(head)
	}

	root.Right = sortedListToBST(slow.Next)
	return root
}

var h *ListNode

var h *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	lenth := 0
	h = head
	for head != nil {
		lenth++
		head = head.Next
	}
	return dfs2(0, lenth-1)
}

func dfs2(l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	left := dfs2(l, mid-1)
	root := &TreeNode{Val: h.Val}
	h = h.Next
	root.Left = left
	root.Right = dfs2(mid+1, r)
	return root
}
