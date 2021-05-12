package leetcode

// https://leetcode-cn.com/problems/missing-number/

func missingNumber(nums []int) int {
	// 方法1
	result := len(nums)
	for i, v := range nums {
		result ^= i ^ v
	}
	return result

	// 方法2
	// var sum int
	// for i, n := range nums {
	// 	sum += (i - n)
	// }
	// return sum + len(nums)
}
