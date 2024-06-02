package main

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	vir := &ListNode{
		Next: head,
	}
	slow, fast := vir, vir

	for slow.Next != nil && slow.Next.Val < x {
		slow = slow.Next
	}

	fast = slow
	for fast.Next != nil {
		if fast.Next.Val >= x {
			fast = fast.Next
			continue
		}
		small := fast.Next
		fast.Next = fast.Next.Next

		tmp := slow.Next
		slow.Next = small
		small.Next = tmp
		slow = slow.Next
	}
	return vir.Next
}

// @lc code=end
