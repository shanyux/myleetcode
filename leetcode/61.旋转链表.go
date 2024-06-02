package main

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	total := 0
	tail := head
	for p := head; p != nil; p = p.Next {
		total++
		tail = p
	}
	tail.Next = head
	k = k % total

	cur := head
	for i := 0; i < total-k-1; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil

	return head
}

// @lc code=end
