package leetcode

// https://leetcode-cn.com/problems/non-decreasing-array/

func checkPossibility(nums []int) bool {
	prev := nums[0]
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < prev {
			nums[i] = prev
			if cnt == 1 {
				return false
			}
			cnt++
		}
		prev = nums[i]
	}

	return true
}
