package leetcode

// https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/

func findLengthOfLCIS(nums []int) int {
	ans := 1
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt++
		} else {
			if cnt > ans {
				ans = cnt
			}
			cnt = 1
		}
	}

	if cnt > ans {
		ans = cnt
	}

	return ans
}
