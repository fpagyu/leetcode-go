package leetcode

// https://leetcode-cn.com/problems/maximum-average-subarray-i/

func findMaxAverage(nums []int, k int) float64 {
	var sum, ans float64

	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	ans = sum
	for i := k; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		if sum > ans {
			ans = sum
		}
	}

	return ans / float64(k)
}
