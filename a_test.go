package leetcode

import "testing"

func Test_Solution(t *testing.T) {
	// r := findMedianSortedArrays([]int{}, []int{1})
	// r := threeSum([]int{})
	// t.Log("result: ", dst)
	// t.Log("result: ", moveZeroes([]int{0, 1, 0, 3, 12}))
	arr := []int{1, 12, -5, -6, 50, 3}
	t.Log("result-0: ", findMaxAverage(arr, 4))
}
