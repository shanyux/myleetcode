/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func maximumGap(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	minV, maxV := math.MaxInt, math.MinInt
	for _, v := range nums {
		if minV > v {
			minV = v
		}
		if maxV < v {
			maxV = v
		}
	}
	if maxV-minV <= 1 {
		return maxV - minV
	}
	fmt.Println(maxV, minV)
	l := len(nums)
	interval := (maxV - minV + l - 2) / (l - 1)
	buckets := ((maxV - minV) / interval) + 1
	minArr, maxArr := make([]int, buckets), make([]int, buckets)
	for i, _ := range minArr {
		minArr[i] = math.MaxInt
	}
	for i, _ := range maxArr {
		maxArr[i] = math.MinInt
	}

	for _, v := range nums {
		i := (v - minV) / interval
		if maxArr[i] < v {
			maxArr[i] = v
		}
		if minArr[i] > v {
			minArr[i] = v
		}
	}
	fmt.Println(minArr, maxArr)
	res := maxArr[0] - minArr[0]
	preV := maxArr[0]
	for i := 1; i < buckets; i++ {
		if maxArr[i] == math.MinInt && minArr[i] == math.MaxInt {
			continue
		}
		if (minArr[i] - preV) > res {
			res = minArr[i] - preV
		}
		preV = maxArr[i]

	}
	return res
}

func main() {
	fmt.Println(maximumGap([]int{100, 3, 2, 1}))
}

// @lc code=end
