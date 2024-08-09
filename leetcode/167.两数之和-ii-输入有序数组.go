/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
package main

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		a, b := numbers[i], numbers[j]
		if a+b == target {
			return []int{i + 1, j + 1}
		}
		if a+b < target {
			i++
		} else {
			j--

		}
	}
	return nil
}

// @lc code=end
