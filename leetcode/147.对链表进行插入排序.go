/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	prev, cur := head, head.Next

	for cur != nil {
		if prev.Val > cur.Val {
			tmp := dummy
			for tmp.Next.Val < cur.Val {
				tmp = tmp.Next
			}
			prev.Next = cur.Next
			cur.Next = tmp.Next
			tmp.Next = cur
			cur = prev.Next
		} else {
			prev = prev.Next
			cur = cur.Next
		}
	}
	return dummy.Next

}

// @lc code=end
