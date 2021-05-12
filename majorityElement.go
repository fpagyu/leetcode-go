package leetcode

// https://leetcode-cn.com/problems/majority-element/

// 时间复杂度O(n), 空间复杂度O(n)
// func majorityElement(nums []int) int {
// 	target := len(nums) / 2
// 	set := make(map[int]int)
// 	for _, n := range nums {
// 		set[n] += 1
// 		if v := set[n]; v > target {
// 			return n
// 		}
// 	}
// 	return 0
// }

// 摩尔投票法
// 从第一个数开始， 如果遇到相同的旧+1，不同的就-1,
// 减到0就重新换个数开始计数, 最后剩下的那个就是众数
func majorityElement(nums []int) int {
	cnt := 1
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == result {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			result = nums[i+1]
		}
	}

	return result
}
