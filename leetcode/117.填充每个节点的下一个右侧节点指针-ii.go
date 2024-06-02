/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	arr := make([]*Node, 0)

	var dfs func(root *Node, depth int)
	dfs = func(root *Node, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, root)
		} else {
			arr[depth].Next = root
			arr[depth] = root
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)

	}
	dfs(root, 0)
	return root
}

// @lc code=end
