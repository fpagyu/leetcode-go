package leetcode

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 删除有序数组中的重复项

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for p := 1; p < len(nums); p++ {
		if nums[i] != nums[p] {
			nums[i+1] = nums[p]
			i++
		}
	}

	return i + 1
}
