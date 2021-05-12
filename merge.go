package leetcode

// https://leetcode-cn.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := 0
	p2 := 0

	for p1 < m && p2 < n {
		if nums1[p1] >= nums2[p2] {
			// 在p1出插入p2处的值
			for i := m; i > p1; i-- {
				nums1[i] = nums1[i-1]
			}
			nums1[p1] = nums2[p2]
			p2++
			m++
		}
		p1++
	}

	for p2 < n {
		nums1[m] = nums2[p2]
		p2++
		m++
	}
}
