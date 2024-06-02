package main

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	res := make([]int, n)
	res[0] = nums[0]
	for i := 1; i < n; i++ {
		res[i] = maxNum2(res[i-1]+nums[i], nums[i])
	}
	r := res[0]
	for i := 1; i < n; i++ {
		r = maxNum2(res[i], r)
	}
	return r

}

//动态规划优化版
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	res := make([]int, n)
	res[0] = nums[0]
	r := res[0]
	for i := 1; i < n; i++ {
		res[i] = maxNum2(res[i-1]+nums[i], nums[i])
		r = maxNum2(res[i], r)
	}
	return r

}

func maxNum2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

//滑动窗口
func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	res := nums[0]
	tmp := res

	for right+1 < len(nums) {
		for tmp < 0 {
			tmp -= nums[left]
			left++
		}
		right++
		tmp += nums[right]
		res = maxNum2(res, tmp)

	}
	return res
}

func maxSubArray4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	res := -1 << 31
	tmp := 0

	for right < len(nums) {
		tmp += nums[right]
		res = maxNum2(res, tmp)
		right++
		for tmp < 0 {
			tmp -= nums[left]
			left++
		}

	}
	return res
}

//前缀
func maxSubArray5(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	preArr := make([]int, n+1)
	preArr[0] = 0
	for i := 1; i <= n; i++ {
		preArr[i] = nums[i-1] + preArr[i-1]
	}

	minVal := 1 >> 21
	maxV := -1 << 31
	for i := 0; i < n; i++ {
		minVal = minNum2(minVal, preArr[i])
		maxV = maxNum2(maxV, preArr[i+1]-minVal)
	}

	return maxV
}

func minNum2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
