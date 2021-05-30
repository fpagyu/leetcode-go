package leetcode

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}

	return nums[0]
}
