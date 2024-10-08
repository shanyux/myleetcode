/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([]int, len(nums)+1)
	tmp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		tmp[i] = nums[i-1] + tmp[i-2]
		if tmp[i] < tmp[i-1] {
			tmp[i] = tmp[i-1]
		}
	}
	return tmp[len(nums)]
}

// @lc code=end

