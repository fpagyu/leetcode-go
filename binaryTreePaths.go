package leetcode

import "fmt"

// https://leetcode-cn.com/problems/binary-tree-paths/

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	l := binaryTreePaths(root.Left)
	r := binaryTreePaths(root.Right)

	if len(r)+len(l) == 0 {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	res := make([]string, 0, len(l)+len(r))

	for i := range l {
		res = append(res, fmt.Sprintf("%d->%s", root.Val, l[i]))
	}

	for i := range r {
		res = append(res, fmt.Sprintf("%d->%s", root.Val, r[i]))
	}

	return res
}
