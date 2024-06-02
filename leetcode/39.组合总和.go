package main

import "fmt"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

func combinationSum1(candidates []int, target int) [][]int {
	reslut := make([][]int, 0)
	tmpres := make([]int, 0)
	var tranck func(tmpres []int, start, target int)
	tranck = func(tmpres []int, start, target int) {
		if target == 0 {
			reslut = append(reslut, append([]int{}, tmpres...))
		}

		if target < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target { // 剪枝，提前返回
				break
			}
			tranck(append(tmpres, candidates[i]), i, target-candidates[i])
		}

		return
	}
	tranck(tmpres, 0, target)
	return reslut
}

func combinationSum(candidates []int, target int) [][]int {
	reslut := make([][]int, 0)
	tmpres := make([]int, 0)
	var tranck func(tmpres []int, start, target int) []int
	tranck = func(tmpres []int, start, target int) []int {
		nums := 0
		for _, n := range tmpres {
			nums += n
		}
		if 0 == target {
			return tmpres
		}

		if target < 0 {
			return nil
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target { // 剪枝，提前返回
				break
			}
			res := tranck(append(tmpres, candidates[i]), i, target-candidates[i])
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

func main() {
	fmt.Print(combinationSum([]int{2, 3, 5}, 8))
}

// @lc code=end
