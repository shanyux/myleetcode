package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	res := &ListNode{}
	tmp := res
	carry := 0
	for {
		if p1 == nil && p2 == nil && carry == 0 {
			break
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		v := 0
		if p1 != nil {
			v = v + p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			v = v + p2.Val
			p2 = p2.Next
		}
		v = v + carry

		tmp.Val = v % 10
		carry = v / 10
	}
	return res.Next
}

// @lc code=end
