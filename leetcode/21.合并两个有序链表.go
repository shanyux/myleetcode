package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{Val: -1, Next: nil}
	p := res
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
			p = p.Next
			continue
		}
		p.Next = p1
		p1 = p1.Next
		p = p.Next
	}
	for p1 != nil {
		p.Next = p1
		p1 = p1.Next
		p = p.Next
	}

	for p2 != nil {
		p.Next = p2
		p2 = p2.Next
		p = p.Next

	}

	return res.Next
}

// @lc code=end
