/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start
package main

import "fmt"

func maxPoints(points [][]int) int {
	if len(points) < 1 {
		return 0
	}
	var res int = 1

	for i := 0; i < len(points); i++ {
		vMmap := make(map[string]int)
		var max int = 0
		for j := i + 1; j < len(points); j++ {
			a, b := points[i][0]-points[j][0], points[i][1]-points[j][1]
			k := getGcd(a, b)
			key := fmt.Sprintf("%d_%d", a/k, b/k)
			if v, ok := vMmap[key]; !ok {
				vMmap[key] = 1
			} else {
				vMmap[key] = v + 1
			}
			if max < vMmap[key] {
				max = vMmap[key]
			}
		}
		if res < max+1 {
			res = max + 1
		}
	}
	return res
}

func getGcd(a, b int) int {
	if b == 0 {
		return a
	}
	return getGcd(b, a%b)
}

// @lc code=end
