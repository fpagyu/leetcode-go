package leetcode

import "sort"

// https://leetcode-cn.com/problems/contains-duplicate/

// 哈希法
// func containsDuplicate(nums []int) bool {
// 	set := make(map[int]struct{})
// 	for _, n := range nums {
// 		if _, has := set[n]; has {
// 			return true
// 		}
// 		set[n] = struct{}{}
// 	}
// 	return false
// }

// 排序判断相连两个元素是否有相同的
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
