package main

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  https://leetcode.cn/problems/linked-list-cycle-ii/solutions/12616/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	f, s := head, head
	for f != nil {
		f = f.Next
		if f == nil {
			return nil
		}
		f = f.Next
		s = s.Next
		if f == s {
			break
		}
	}

	if f == nil {
		return nil
	}

	f = head
	for f != s {
		f = f.Next
		s = s.Next
	}
	return f

}

// @lc code=end
