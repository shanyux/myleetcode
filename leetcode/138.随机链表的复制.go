package main

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 随机链表的复制
 */

// @lc code=start

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		n := &Node{
			Val: p.Val,
		}
		n.Next = p.Next
		p.Next = n
		p = n.Next
	}

	p = head
	for p != nil {
		if p.Random == nil {
			p = p.Next.Next
			continue
		}
		p.Next.Random = p.Random.Next
		p = p.Next.Next
	}

	p = head
	dummy := &Node{}
	cur := dummy
	for p != nil {
		cur.Next = p.Next
		cur = cur.Next
		p.Next = cur.Next
		p = p.Next
	}

	return dummy.Next
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)

	p := head
	for p != nil {
		n := &Node{
			Val: p.Val,
		}
		nodeMap[p] = n
		p = p.Next
	}

	p = head
	dummy := &Node{}
	cur := dummy
	for p != nil {
		cur.Next = nodeMap[p]
		cur.Next.Random = nodeMap[p.Random]
		cur = cur.Next
		p = p.Next
	}

	return dummy.Next
}

// @lc code=end
