package main

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	f, s := head, head

	for f != nil && f.Next != nil && f.Next.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	tail := s.Next
	s.Next = nil
	tail = reverseList(tail)

	for tail != nil {
		newTail := tail.Next
		tail.Next = head.Next
		head.Next = tail
		head = tail.Next
		tail = newTail
	}

}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	tail := head
	for tail != nil {
		newHead := tail.Next
		tail.Next = prev
		prev = tail
		tail = newHead
	}
	return prev
}

// @lc code=end
