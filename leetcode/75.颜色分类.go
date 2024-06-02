package main

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start

// （left,right）是1的范围
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	cur := left
	for cur < right {
		if nums[cur] == 1 {
			cur++
			continue
		}
		if nums[cur] == 0 {
			swap(nums, left, cur)
			left++
			cur++
		} else {
			swap(nums, right, cur)
			right--

		}

	}

}

func swap(nums []int, a, b int) {
	c := nums[b]
	nums[b] = nums[a]
	nums[a] = c
}

// @lc code=end
