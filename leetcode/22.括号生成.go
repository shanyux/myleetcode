package main

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	left, right := n, n
	tmpres := []byte{}
	trackRes(left, right, &tmpres, &res)
	return res
}

func trackRes(left, right int, tmpres *[]byte, res *[]string) {
	if left < 0 || right < 0 {
		return
	}
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, string(*tmpres))
		return
	}

	// l := len(*tmpres)
	// (*tmpres)[l] = '('
	*tmpres = append(*tmpres, '(')
	trackRes(left-1, right, tmpres, res)
	*tmpres = (*tmpres)[:len(*tmpres)-1]

	*tmpres = append(*tmpres, ')')
	trackRes(left, right-1, tmpres, res)
	*tmpres = (*tmpres)[:len(*tmpres)-1]
}

// @lc code=end
