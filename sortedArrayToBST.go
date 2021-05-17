package leetcode

// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	pivot := len(nums) / 2
	root := &TreeNode{Val: nums[pivot]}

	if pivot > 0 {
		root.Left = sortedArrayToBST(nums[:pivot])
	}

	if i := pivot + 1; i < len(nums) {
		root.Right = sortedArrayToBST(nums[i:])
	}

	return root
}
