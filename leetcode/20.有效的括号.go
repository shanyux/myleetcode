package main

import "fmt"

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

type Stack []byte

func (s *Stack) pop() byte {
	if len(*s) == 0 {
		// fmt.Println(3)
		return 0
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) push(c byte) {
	*s = append(*s, c)
}

func isValid(s string) bool {
	stack := &Stack{}
	for i := 0; i < len(s); i++ {
		if validLeft(s[i]) == ' ' {
			stack.push(s[i])
			continue
		}
		if len(*stack) == 0 || validLeft(s[i]) != stack.pop() {

			return false
		}
	}

	return len(*stack) == 0
}

func validLeft(c byte) byte {
	if c == ')' {
		return '('
	}
	if c == '}' {
		return '{'
	}
	if c == ']' {
		return '['
	}
	return ' '
}

func main() {
	fmt.Println(isValid("()"))
}

// @lc code=end
