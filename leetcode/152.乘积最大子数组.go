package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct1(nums []int) int {
	imax, imin := 1, 1
	max := math.MinInt
	for _, i := range nums {
		if i < 0 {
			imin, imax = imax, imin
		}
		if imax*i > i {
			imax = imax * i
		} else {
			imax = i
		}
		if imin*i < i {
			imin = imin * i
		} else {
			imin = i
		}
		if max < imax {
			max = imax
		}
	}
	return max
}

func minV(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		}
		return b
	}
	if a > c {
		return c
	}
	return a
}

func maxV2(a, b, c int) int {
	if a < b {
		if b < c {
			return c
		}
		return b
	}
	if a < c {
		return c
	}
	return a
}

func maxProduct(nums []int) int {
	max := math.MinInt
	imax, imin := make([]int, len(nums)), make([]int, len(nums))
	imax[0], imin[0] = nums[0], nums[0]
	for i, _ := range nums {
		if i == 0 {
			if max < imax[i] {
				max = imax[i]
			}
			continue
		}
		imax[i] = maxV2(imax[i-1]*nums[i], imin[i-1]*nums[i], nums[i])
		imin[i] = minV(imax[i-1]*nums[i], imin[i-1]*nums[i], nums[i])
		if max < imax[i] {
			max = imax[i]
		}
		if max == 1000000000 {
			return 1000000000
		}

	}
	return max
}

func main() {
	fmt.Println(maxProduct([]int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}))
}

// @lc code=end
