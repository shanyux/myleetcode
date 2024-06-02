package main

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast, count := 0, 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else {
			if slow < fast && count < 1 {
				slow++
				nums[slow] = nums[fast]
				count++
			}
		}
		fast++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			count = 0
		}
	}
	return slow + 1
}

func removeDuplicate2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := 0
	for _, v := range nums {
		if nums[last-2] != v || last < 2 {
			nums[last] = v
			last++
		}
	}
	return last
}

// @lc code=end
