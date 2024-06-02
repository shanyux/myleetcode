package main

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
//单调栈和哨兵
func largestRectangleArea(heights []int) int {
	arr := make([]int, 0)
	res := 0
	heightsCopy := make([]int, 1) //哨兵
	heightsCopy = append(heightsCopy, heights...)
	heightsCopy = append(heightsCopy, 0) //哨兵
	for i, v := range heightsCopy {
		for len(arr) > 0 && heightsCopy[arr[len(arr)-1]] > v {
			cur := arr[len(arr)-1]
			arr = arr[:len(arr)-1]

			l := i - (arr[len(arr)-1] + 1) //栈顶后面那个  可能被弹出去了
			h := heightsCopy[cur]
			res = maxV(res, l*h)

		}
		arr = append(arr, i)
	}
	return res
}

func maxV(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
