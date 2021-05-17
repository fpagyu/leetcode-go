package leetcode

// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 快慢指针的方式确定根节点的位置
// func sortedListToBST(head *ListNode) *TreeNode {
// 	if head == nil {
// 		return nil
// 	}
// 	var prev *ListNode
// 	slow := head
// 	fast := head
// 	for fast != nil && fast.Next != nil {
// 		prev = slow
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	root := &TreeNode{Val: slow.Val}
// 	if prev != nil {
// 		prev.Next = nil
// 		root.Left = sortedListToBST(head)
// 	}
// 	root.Right = sortedListToBST(slow.Next)
// 	return root
// }

// 使用中序遍历+分治法
// TODO: 重刷
func sortedListToBST(head *ListNode) *TreeNode {
	return nil
}
