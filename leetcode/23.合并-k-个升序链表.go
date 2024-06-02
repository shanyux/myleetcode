package main

import "container/heap"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	res := &ListNode{Val: -1}
	tmp := res
	queue := make(Queue, 0, len(lists))
	for _, l := range lists {
		if l == nil {
			continue
		}
		queue = append(queue, l)
	}
	// for i, head := range lists {
	// 	queue[i] = head
	// }
	heap.Init(&queue)

	for queue.Len() != 0 {
		node := heap.Pop(&queue).(*ListNode)
		tmp.Next = node
		if node.Next != nil {
			heap.Push(&queue, node.Next)
		}
		tmp = tmp.Next
	}
	return res.Next
}

type Queue []*ListNode

func (q *Queue) Len() int           { return len(*q) }
func (q Queue) Less(i, j int) bool  { return q[i].Val < q[j].Val }
func (q Queue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *Queue) Push(x interface{}) { *q = append(*q, x.(*ListNode)) }
func (q *Queue) Pop() interface{} {
	res := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return res
}

// @lc code=end
// @lc code=end
// @lc code=end
// @lc code=end
