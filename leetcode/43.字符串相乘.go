package main

import "fmt"

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	res := make([]int, len1+len2)

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			p1, p2 := i+j, i+j+1
			left := int(num1[i] - '0')
			right := int(num2[j] - '0')
			tmp := left*right + res[p2]
			res[p2] = tmp % 10
			res[p1] += tmp / 10
		}
	}

	fmt.Println(res)
	strres := ""
	i := 0
	for ; i < len(res); i++ {
		if res[i] == 0 {
			continue
		}
		break
	}
	for ; i < len(res); i++ {
		strres += string(rune(res[i] + '0'))
	}

	if strres == "" {
		return "0"
	}
	return strres
}

func main() {
	fmt.Println(multiply("123", "456"))
}

// @lc code=end
