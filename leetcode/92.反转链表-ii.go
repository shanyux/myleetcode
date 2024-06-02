package main

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//https://leetcode.cn/problems/reverse-linked-list-ii/solutions/37247/bu-bu-chai-jie-ru-he-di-gui-di-fan-zhuan-lian-biao/
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	var succ *ListNode
	var revertN func(head *ListNode, n int) *ListNode

	revertN = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			succ = head.Next
			return head
		}
		last := revertN(head.Next, n-1)
		head.Next.Next = head
		head.Next = succ
		return last
	}
	if left == 1 {
		return revertN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// @lc code=end
