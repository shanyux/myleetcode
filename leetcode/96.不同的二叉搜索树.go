package opsappservice

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int { //构造树会超时
	if n == 0 {
		return nil
	}

	res := build(1, n)
	return len(res)

}

func build(lo, hi int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for i := lo; i <= hi; i++ {
		left := build(lo, i-1)
		right := build(i+1, hi)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}

func numTrees(n int) int { //构造树
	if n == 0 {
		return 0
	}
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ { //几个节点
		for j := 0; j < i; j++ { //左边几个
			res[i] += res[j] * res[i-j-1]
		}
	}
	return res[n]
}

// @lc code=end
