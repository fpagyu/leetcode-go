package leetcode

import "fmt"

// https://leetcode-cn.com/problems/summary-ranges/

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	result := make([]string, 0)

	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if i-1 == p {
				result = append(result, fmt.Sprintf("%d", nums[p]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[p], nums[i-1]))
			}
			p = i
		}
	}

	if p == len(nums)-1 {
		result = append(result, fmt.Sprintf("%d", nums[p]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[p], nums[len(nums)-1]))
	}

	return result
}
