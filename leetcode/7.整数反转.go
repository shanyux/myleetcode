package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	00

	tmp := x

	res := 0
	for tmp != 0 {
		if t := int32(res); t*10/10 != t {
			return 0
		}
		res = res*10 + tmp%10
		tmp = tmp / 10
	}
	// res := 0
	// for _, i := range array {
	// 	if t := int32(res); t*10/10 != t {
	// 		return 0
	// 	}
	// 	res = res*10 + i

	// }
	return res
}

func main() {
	fmt.Print(reverse(-156469))
}

// @lc code=end
