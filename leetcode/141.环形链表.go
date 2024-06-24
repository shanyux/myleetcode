package main

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	f, s := head, head
	f = f.Next.Next
	s = s.Next
	for f != nil {
		if f == s {
			return true
		}
		if f.Next == nil {
			return false
		}
		f = f.Next.Next
		s = s.Next
	}
	return false
}

// @lc code=end
