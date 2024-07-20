/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
package main

type childStruct struct {
	cur int
	min int
}
type MinStack struct {
	childStruct []*childStruct
}

func Constructor() MinStack {
	return MinStack{
		childStruct: make([]*childStruct, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.childStruct) == 0 {
		this.childStruct = append(this.childStruct, &childStruct{
			cur: val,
			min: val,
		})
		return
	}
	minV := this.childStruct[len(this.childStruct)-1].min
	if minV > val {
		minV = val
	}
	this.childStruct = append(this.childStruct, &childStruct{
		cur: val,
		min: minV,
	})
}

func (this *MinStack) Pop() {
	this.childStruct = this.childStruct[:len(this.childStruct)-1]
}

func (this *MinStack) Top() int {
	return this.childStruct[len(this.childStruct)-1].cur
}

func (this *MinStack) GetMin() int {
	return this.childStruct[len(this.childStruct)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
