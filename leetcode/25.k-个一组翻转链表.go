package main

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return head
	}
	head1, cur := head, head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	res := reverseN(head, cur)
	head1.Next = reverseKGroup(cur, k)
	return res

}

func reverseN(head *ListNode, end *ListNode) *ListNode {
	pre, cur, nex := (*ListNode)(nil), head, head
	for cur != end {
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}

// @lc code=end
