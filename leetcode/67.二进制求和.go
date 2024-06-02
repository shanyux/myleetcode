package main

import "fmt"

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	len1, len2 := len(a)-1, len(b)-1
	reslen := len1 + 1
	if reslen < len2+1 {
		reslen = len2 + 1
	}
	res := make([]int, reslen+1)
	carray := 0
	index := 0
	for len1 >= 0 && len2 >= 0 {
		l, r := int(a[len1]-'0'), int(b[len2]-'0')
		value := l + r + carray
		carray = value / 2
		res[index] = value % 2
		len1--
		len2--
		index++
	}

	for len1 >= 0 {
		l := int(a[len1] - '0')
		value := l + carray
		carray = value / 2
		res[index] = value % 2
		len1--
		index++

	}

	for len2 >= 0 {
		r := int(b[len2] - '0')
		value := r + carray
		carray = value / 2
		res[index] = value % 2
		len2--
		index++
	}
	if carray > 0 {
		res[index] = carray
	}
	begin := true
	res1 := make([]byte, 0)
	fmt.Println(res)

	for right := len(res) - 1; right >= 0; right-- {
		if begin && res[right] == 0 {
			continue
		}
		begin = false
		res1 = append(res1, byte(res[right]+'0'))
	}
	if len(res1) == 0 {
		return "0"
	}
	return string(res1)
}

func main() {
	fmt.Print((addBinary("1", "11")))
}

// @lc code=end
