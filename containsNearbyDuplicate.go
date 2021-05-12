package leetcode

// https://leetcode-cn.com/problems/contains-duplicate-ii/

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	p := 0
// 	for p < len(nums)-1 {
// 		for i := p + 1; i <= p+k && i < len(nums); i++ {
// 			if nums[p] == nums[i] {
// 				return true
// 			}
// 		}
// 		p++
// 	}
// 	return false
// }

func containsNearbyDuplicate(nums []int, k int) bool {
	if k > len(nums) {
		k = len(nums)
	}

	set := make(map[int]int)
	for i := 0; i < k; i++ {
		set[nums[i]] += 1
		if set[nums[i]] > 1 {
			return true
		}
	}

	for i := k; i < len(nums); i++ {
		set[nums[i]] += 1
		if set[nums[i]] > 1 {
			return true
		}
		set[nums[i-k]] -= 1
	}

	return false
}
