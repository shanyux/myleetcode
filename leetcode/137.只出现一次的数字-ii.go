package main

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	count := make([]int, 32)
	for _, v := range nums {

		for i := 0; i < 32; i++ {
			count[i] += 1 & v
			v >>= 1
		}
	}
	// fmt.Println(count)
	res := 0
	m := 3
	for i := 31; i >= 0; i-- {
		res = res << 1
		res |= count[i] % m

	}
	if count[31]%m == 0 {
		return res
	} else {
		return int(int32(res))
	}
}

func main() {
	singleNumber([]int{2, 2, 3, 2})
}

// @lc code=end
