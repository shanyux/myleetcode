package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)

	var backtrack func(start int, tmp []int)
	backtrack = func(start int, tmp []int) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)
		for i := start; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			backtrack(i+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrack(0, tmp)
	return res
}

// @lc code=end
