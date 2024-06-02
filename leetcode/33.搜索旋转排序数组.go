package main

import "fmt"

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if target == nums[mid] {
			return mid
		}
		fmt.Println(lo, mid, hi)
		if nums[lo] <= nums[mid] { //前半段
			fmt.Println(lo, mid, hi)
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
			continue
		}
		if target <= nums[hi] && target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
		fmt.Println(lo, mid, hi)

	}
	return -1
}

func main() {
	fmt.Print(search([]int{3, 1}, 1))
}

// @lc code=end
