package leetcode

// https://leetcode-cn.com/problems/max-consecutive-ones/

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0

	for _, n := range nums {
		if n == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}

	return max
}
