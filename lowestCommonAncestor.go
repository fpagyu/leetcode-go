package leetcode

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}

	if root.Val >= q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val <= p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return nil
}
