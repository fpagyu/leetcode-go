package leetcode

// 方法1
// func moveZeroes(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			for j := i + 1; j < len(nums); j++ {
// 				if nums[j] != 0 {
// 					nums[i], nums[j] = nums[j], nums[i]
// 					break
// 				}
// 			}
// 		}
// 	}
// }

func moveZeroes(nums []int) {
	p := 0

	for _, n := range nums {
		if n != 0 {
			nums[p] = n
			p++
		}
	}
	for p < len(nums) {
		nums[p] = 0
		p++
	}
}
