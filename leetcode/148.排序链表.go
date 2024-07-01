package main

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//slow就是中间
	mid := slow.Next
	slow.Next = nil
	left, right := sortList(head), sortList(mid)

	dummy := &ListNode{
		// Next: head,
	}
	res := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			res.Next = left
			left, res = left.Next, res.Next
		} else {
			res.Next = right
			right, res = right.Next, res.Next
		}
	}
	if left != nil {
		res.Next = left

	}
	if right != nil {
		res.Next = right

	}
	return dummy.Next
}

// @lc code=end
