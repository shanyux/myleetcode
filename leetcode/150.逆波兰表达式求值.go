/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
package main

import "strconv"

func evalRPN(tokens []string) int {
	arr := make([]int, 0)
	for _, s := range tokens {
		if i, err := strconv.Atoi(s); err == nil {
			arr = append(arr, i)
		} else {
			v1, v2 := arr[len(arr)-2], arr[len(arr)-1]
			arr = arr[:len(arr)-2]
			switch s {
			case "+":
				arr = append(arr, v1+v2)
			case "-":
				arr = append(arr, v1-v2)
			case "*":
				arr = append(arr, v1*v2)
			case "/":
				arr = append(arr, v1/v2)
			}
		}
	}
	return arr[0]
}

// @lc code=end
