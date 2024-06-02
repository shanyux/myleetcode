package main

import "fmt"

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	isUse := make([]bool, len(nums))
	var trackRes func(tmp []int)
	trackRes = func(tmp []int) {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}
		// fmt.Println(tmp)
		for i := 0; i < len(nums); i++ {
			if isUse[i] {
				continue
			}
			isUse[i] = true
			trackRes(append(tmp, nums[i]))
			isUse[i] = false
		}

	}
	tmp := make([]int, 0)
	trackRes(tmp)
	return res

}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

// @lc code=end
