package leetcode

import "sort"

// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
// 重刷
func maximumProduct(nums []int) int {
	sort.Ints(nums)

	a := nums[0] * nums[1] * nums[len(nums)-1]
	b := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	if a > b {
		return a
	}
	return b
}

// func maximumProduct(nums []int) int {
// 	return 0
// }
