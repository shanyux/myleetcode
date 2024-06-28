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
package main

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

func reorderList(head *ListNode) {
	len := 0
	h := head
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	for h != nil {
		h = h.Next
		len++
	}
	reorderList1(head, len)
}

func reorderList1(head *ListNode, len int) *ListNode {
	if len == 1 {
		res := head.Next
		head.Next = nil
		return res
	}

	if len == 2 {
		res := head.Next.Next
		head.Next.Next = nil
		return res
	}

	tail := reorderList1(head.Next, len-2)
	hn := head.Next
	head.Next = tail
	tn := tail.Next
	tail.Next = hn
	return tn
}

// @lc code=end
