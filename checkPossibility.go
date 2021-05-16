package leetcode

// https://leetcode-cn.com/problems/non-decreasing-array/

// todo: 重刷
func checkPossibility(nums []int) bool {
	min := -100000

	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if min > nums[i+1] {
				nums[i+1] = nums[i]
			} else {
				nums[i] = min
			}
			cnt++
		}

		if cnt > 1 {
			return false
		}
		min = nums[i]
	}

	return true
}
