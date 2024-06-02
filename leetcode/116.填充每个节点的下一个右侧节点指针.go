/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
package main

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	pre := root
	for pre.Left != nil {
		tmp := pre
		for tmp != nil {
			tmp.Left.Next = tmp.Right
			if tmp.Next != nil {
				tmp.Right.Next = tmp.Next.Left
			}
			tmp = tmp.Next
		}
		pre = pre.Left
	}
	return root
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		tmp := queue[0]
		if tmp.Left != nil {
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil {
			queue = append(queue, tmp.Right)
		}
		for i := 1; i < l; i++ {
			q := queue[i]
			tmp.Next = q
			tmp = tmp.Next

			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[l:]
	}
	return root
}

// @lc code=end
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	cur := root
	for cur != nil {
		pre := &Node{}
		tmp := pre
		// cur := pre
		for cur != nil {
			if cur.Left != nil {
				tmp.Next = cur.Left
				tmp = tmp.Next
			}
			if cur.Right != nil {
				tmp.Next = cur.Right
				tmp = tmp.Next
			}
			cur = cur.Next
		}
		cur = pre.Next
	}
	return root
}
