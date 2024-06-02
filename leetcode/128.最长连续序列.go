/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
package main

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]int)
	maxLen := 0
	//像毒圈一样  一点一点蚕食扩大范围   只记录连续数字的边界的最大长度
	for _, v := range nums {
		if numsMap[v] > 0 { //已经存在的不用再记
			continue
		}
		left := numsMap[v-1]
		right := numsMap[v+1]
		curLen := left + right + 1
		if curLen > maxLen {
			maxLen = curLen
		}
		numsMap[v] = curLen
		numsMap[v-left] = curLen
		numsMap[v+right] = curLen

	}

	return maxLen

}

// @lc code=end
