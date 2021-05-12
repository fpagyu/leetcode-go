package leetcode

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1 := 0
	p2 := 0

	var v0, v1 int
	length := len(nums1) + len(nums2)
	for i := 0; i < (length/2 + 1); i++ {
		v0 = v1

		if p1 == len(nums1) {
			v1 = nums2[p2]
			p2++
			continue
		}

		if p2 == len(nums2) {
			v1 = nums1[p1]
			p1++
			continue
		}

		if nums1[p1] < nums2[p2] {
			v1 = nums1[p1]
			p1++
		} else {
			v1 = nums2[p2]
			p2++
		}
	}

	if length%2 == 0 {
		return float64(v0+v1) / 2.0
	}

	return float64(v1)
}
