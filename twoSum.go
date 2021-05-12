package leetcode

// https://leetcode-cn.com/problems/two-sum/
// 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if index, ok := m[target-nums[i]]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}

	return nil
}
