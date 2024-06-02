/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start

//  * Definition for a Node.
package main

// type Node struct {
// 	Val       int
// 	Neighbors []*Node
// }

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) != 0 {
		l := len(queue)
		for _, q := range queue[:l] {
			node := &Node{
				Val: q.Val,
			}
			if len(q.Neighbors) != 0 {
				queue = append(queue, q.Neighbors...)
			}
		}
		queue = queue[l:]
	}

}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, node)
	nodeMap := make(map[int]*Node)
	n := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	nodeMap[node.Val] = n

	for len(queue) > 0 {
		l := len(queue)
		for _, qqqqqq := range queue[:l] {
			for _, ne := range qqqqqq.Neighbors {
				_, ok := nodeMap[ne.Val]
				if !ok {
					n := &Node{
						Val:       ne.Val,
						Neighbors: make([]*Node, 0),
					}
					nodeMap[ne.Val] = n
					queue = append(queue, ne)

				}
				nodeMap[qqqqqq.Val].Neighbors = append(nodeMap[qqqqqq.Val].Neighbors, nodeMap[ne.Val])

			}
		}
		queue = queue[l:]
	}
	return n
}

func cloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodeMap := make(map[int]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		c, ok := nodeMap[node.Val]
		if ok {
			return c
		}
		n := &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, 0),
		}
		nodeMap[node.Val] = n
		for _, ne := range node.Neighbors {
			n.Neighbors = append(n.Neighbors, dfs(ne))
		}

		return n
	}

	return dfs(node)
}

// @lc code=end
