package leetcode

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		nums[(num-1)%len(nums)] += len(nums)
	}

	ret := make([]int, 0)
	for i, n := range nums {
		if n <= len(nums) {
			ret = append(ret, i+1)
		}
	}
	return ret
}
