package main

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	delta := math.MaxInt
	num := 0
	fmt.Println(nums)
	for i, v := range nums {
		if i >= len(nums)-2 {
			break
		}
		num = v + twoSumClosest(nums[i+1:], target-v)

		if math.Abs(float64(target-num)) < math.Abs(float64(delta)) {
			delta = target - num
			fmt.Println(delta)
		}
	}
	return target - delta
}

func twoSumClosest(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	delta := math.MaxInt
	num := 0

	for lo < hi {
		num = nums[lo] + nums[hi]
		if math.Abs(float64(target-num)) < math.Abs(float64(delta)) {
			delta = target - num
		}
		if target > num {
			lo++
		} else {
			hi--
		}
	}
	return target - delta
}

func main() {
	fmt.Println(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
}

// @lc code=end
