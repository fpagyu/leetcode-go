package leetcode

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeBuilder(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	q := list.New()
	root := &TreeNode{Val: arr[0].(int)}
	q.PushBack(root)
	for i := 1; i < len(arr); i += 2 {
		e := q.Front()
		q.Remove(e)
		p := e.Value.(*TreeNode)

		if v, ok := arr[i].(int); ok {
			p.Left = &TreeNode{Val: v}
			q.PushBack(p.Left)
		}

		if i+1 < len(arr) {
			if v, ok := arr[i+1].(int); ok {
				p.Right = &TreeNode{Val: v}
				q.PushBack(p.Right)
			}
		}
	}

	return root
}
