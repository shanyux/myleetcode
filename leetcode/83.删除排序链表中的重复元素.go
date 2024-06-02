package main

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := head.Val
	cur := head.Next
	pre := head
	pre.Next = nil
	for cur != nil {
		if cur.Val == last {
			cur = cur.Next
			continue
		}
		last = cur.Val
		pre.Next = cur
		cur = cur.Next

		pre = pre.Next
		pre.Next = nil
	}

	return head

}

//快慢指针
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil

	return head

}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head

}

//递归
func deleteDuplicates5(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
	} else {
		cur := head
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		return deleteDuplicates(cur)
	}

	return head
}

// @lc code=end
