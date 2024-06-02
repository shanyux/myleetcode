package main

import "fmt"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] N 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || numRows >= n {
		return s
	}
	t := 2*numRows - 2 //每个周期包含的字符数量
	c := numRows - 1   //每个周期包含的列数

	totalc := (n + t - 1) / t * c //总共列数
	tmpArrays := make([][]byte, numRows)
	for i, _ := range tmpArrays {
		tmpArrays[i] = make([]byte, totalc)
	}

	x, y := 0, 0
	for i, c := range s {
		tmpArrays[x][y] = byte(c)
		if i%t < numRows-1 { //小于一半往下
			x++
		} else {
			x--
			y++
		}
	}
	res := make([]byte, 0, len(s))
	for _, cos := range tmpArrays {
		for _, c := range cos {
			if c > 0 {
				res = append(res, c)
			}
		}
	}
	return string(res)

}

// @lc code=end

func main() {
	fmt.Print(convert("AB", 1))
}
