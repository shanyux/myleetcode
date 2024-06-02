package main

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	reslut := make([][]int, 0)
	tmpres := make([]int, 0)
	var tranck func(tmpres []int, start, target int)
	tranck = func(tmpres []int, start, target int) {
		if 0 == target {
			reslut = append(reslut, append([]int{}, tmpres...))
		}

		if 0 > target {
			return
		}

		if start >= len(candidates) {
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target { // 剪枝，提前返回
				break
			}
			tranck(append(tmpres, candidates[i]), i+1, target-candidates[i])
			for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
				i++
			}
		}

		return
	}
	tranck(tmpres, 0, target)
	return reslut
}

func combinationSum21(candidates []int, target int) [][]int {
	reslut := make([][]int, 0)
	tmpres := make([]int, 0)
	var tranck func(tmpres []int, start, target int) []int
	tranck = func(tmpres []int, start, target int) []int {
		nums := 0
		for _, n := range tmpres {
			nums += n
		}
		if nums == target {
			return tmpres
		}

		if nums > target {
			return nil
		}

		for i := start; i < len(candidates); i++ {
			res := tranck(append(tmpres, candidates[i]), i+1, target)
			if res != nil {
				// fmt.Println(res)
				t := append([]int{}, res...)
				reslut = append(reslut, t)
			}
		}

		return nil
	}
	tranck(tmpres, 0, target)
	return reslut
}

// @lc code=end
