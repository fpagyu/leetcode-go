package leetcode

// https://leetcode-cn.com/problems/symmetric-tree/
// 给定一个二叉树， 检查它是否镜像对称

func isSymmetric(root *TreeNode) bool {
	return compare(root, root)
}

// 递归方式
func compare(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val &&
		compare(a.Left, b.Right) &&
		compare(a.Right, b.Left)
}
