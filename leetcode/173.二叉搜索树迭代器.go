package main

/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
 */

// @lc code=start

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// type BSTIterator struct {
// 	List []*TreeNode
// }

// func Constructor(root *TreeNode) BSTIterator {
// 	res := &BSTIterator{
// 		List: make([]*TreeNode, 0),
// 	}
// 	inOrder(res, root)
// 	return *res
// }

// func inOrder(bSTIterator *BSTIterator, root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	inOrder(bSTIterator, root.Left)
// 	bSTIterator.List = append(bSTIterator.List, root)
// 	inOrder(bSTIterator, root.Right)
// }

// func (this *BSTIterator) Next() int {
// 	if len(this.List) == 0 {
// 		return 0
// 	}
// 	tmp := this.List[0]
// 	this.List = this.List[1:]
// 	return tmp.Val
// }

// func (this *BSTIterator) HasNext() bool {
// 	return len(this.List) != 0
// }

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

type BSTIterator struct {
	List []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	res := &BSTIterator{
		List: make([]*TreeNode, 0),
	}
	pushStack(res, root)
	return *res
}

func pushStack(bSTIterator *BSTIterator, root *TreeNode) {
	if root == nil {
		return
	}
	bSTIterator.List = append(bSTIterator.List, root)
	root = root.Left
	pushStack(bSTIterator, root)
}

func (this *BSTIterator) Next() int {
	if len(this.List) == 0 {
		return 0
	}
	tmp := this.List[len(this.List)-1]
	this.List = this.List[:len(this.List)-1]
	pushStack(this, tmp.Right)

	return tmp.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.List) != 0
}
