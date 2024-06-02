package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return getNSum1(nums, 4, 0, target)
}

// @lc code=end

func getNSum1(nums []int, n, start, target int) [][]int {
	var res [][]int
	if n < 2 || start >= len(nums) {
		return nil
	}

	if n == 2 {
		lo, hi := start, len(nums)-1
		for lo < hi {
			s := nums[lo] + nums[hi]
			if s < target {
				lo++
				for lo < hi && nums[lo-1] == nums[lo] {
					lo++
				}
			} else if s > target {
				hi--
				for lo < hi && nums[hi+1] == nums[hi] {
					hi--
				}
			} else {
				res = append(res, []int{nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[lo-1] == nums[lo] {
					lo++
				}
				for lo < hi && nums[hi+1] == nums[hi] {
					hi--
				}
			}
		}
	} else {
		for i := start; i < len(nums); i++ {
			tmpRes := getNSum1(nums, n-1, i+1, target-nums[i])
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
	fmt.Print(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}
