package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	valueIndeMap := make(map[int]int, len(nums))
	for i, n := range nums {
		need := target - n
		if _, ok := valueIndeMap[need]; ok {
			return []int{valueIndeMap[need], i}
		}
		valueIndeMap[n] = i
	}

	return nil
}

func main() {
	fmt.Print(twoSum([]int{2, 9, 8, 7, 11, 15}, 9))
}

// @lc code=end
