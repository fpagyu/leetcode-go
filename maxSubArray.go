package leetcode

// https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	var sum int
	max := nums[0]
	for i := range nums {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
