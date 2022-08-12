package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  后序遍历
// 执行用时：0 ms, 在所有 Go 提交中击败了 100% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 18.74% 的用户
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ret []int
	ret = append(ret, postorderTraversal(root.Left)...)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)

	return ret
}
