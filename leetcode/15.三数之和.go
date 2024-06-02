package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return getNSum(nums, 3, 0, 0)
}

func getNSum(nums []int, n, start, target int) [][]int {
	var res [][]int
	if n < 2 || len(nums) < n {
		return nil
	}

	if n == 2 {
		lo, hi := start, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				// lo++
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				// hi--
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{nums[lo], nums[hi]})
				// lo++
				// hi--
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		// for i, v := range nums[start:] {
		// 	if i < len(nums)-1 && v == nums[i+1] {
		// 		continue
		// 	}
		// 	tmpRes := getNSum(nums, n-1, i+1, target-v)
		// 	for _, t := range tmpRes {
		// 		t = append(t, v)
		// 		res = append(res, t)
		// 	}
		// }
		for i := start; i < len(nums); i++ {
			tmpRes := getNSum(nums, n-1, i+1, target-nums[i])
			for _, t := range tmpRes {
				t = append(t, nums[i])
				res = append(res, t)
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}

func main() {
	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

// @lc code=end
