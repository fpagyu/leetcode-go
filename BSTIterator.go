package leetcode

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	*list.List
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		List: list.New(),
	}

	iter.PushBack(root)

	p := root
	for p.Left != nil {
		iter.PushBack(p.Left)
		p = p.Left
	}

	return iter
}

func (this *BSTIterator) Next() int {
	e := this.Back()
	this.Remove(e)
	node := e.Value.(*TreeNode)

	p := node.Right
	for p != nil {
		this.PushBack(p)
		p = p.Left
	}

	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.Len() > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
