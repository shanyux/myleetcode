package main

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := &ListNode{Next: head}
	pre, cur := h, h.Next
	deleteDup := false
	for cur != nil {
		if cur.Next == nil {
			pre.Next = cur
			break
		}
		last := cur.Val
		for cur.Next != nil && cur.Next.Val == last {
			cur = cur.Next
			deleteDup = true
		}
		if deleteDup {
			cur = cur.Next
			deleteDup = false
			pre.Next = nil
			continue
		}
		pre.Next = cur
		pre = pre.Next
		deleteDup = false
		cur = cur.Next
	}

	return h.Next
}

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
	} else {
		next := head.Next
		for next != nil && next.Val == head.Val {
			next = next.Next
		}
		return deleteDuplicates(next)
	}
	return head
}
