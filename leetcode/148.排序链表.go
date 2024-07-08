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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	len, h := 0, head
	for h != nil {
		h = h.Next
		len++
	}
	dummy := &ListNode{
		Next: head,
	}
	w := 1
	for w < len {
		pre, h := dummy, dummy.Next
		for h != nil {
			h1, w1 := h, w
			for h != nil && w1 > 0 {
				h = h.Next
				w1--
			}
			if w1 > 0 {
				break
			}
			h2, w2 := h, w
			for h != nil && w2 > 0 {
				h = h.Next
				w2--
			}
			c1, c2 := w, w-w2
			for c1 > 0 && c2 > 0 {
				if h1.Val > h2.Val {
					pre.Next, h2 = h2, h2.Next
					c2--
				} else {
					pre.Next, h1 = h1, h1.Next
					c1--
				}
				pre = pre.Next
			}
			if c1 > 0 {
				pre.Next = h1
			} else {
				pre.Next = h2
			}

			for c1 > 0 || c2 > 0 {
				pre = pre.Next
				c1--
				c2--
			}
			pre.Next = h
		}
		w *= 2
	}

	return dummy.Next
}
