package main

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge1(nums1 []int, m int, nums2 []int, n int) {
	all := m + n
	for i := m - 1; i >= 0; i-- {
		nums1[all+i-m] = nums1[i]
	}
	i, j, k := 0, n, 0
	for i < n && j < all {
		if nums2[i] >= nums1[j] {
			nums1[k] = nums1[j]
			j++
		} else {
			nums1[k] = nums2[i]
			i++
		}
		k++

	}

	for i < n {
		nums1[k] = nums2[i]
		k++
		i++

	}
}

func main() {
	merge1([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

// @lc code=end
