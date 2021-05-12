package leetcode

// https://leetcode-cn.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}

	return len(nums)
}
