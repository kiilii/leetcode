package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.7 MB, 在所有 Go 提交中击败了 99.84% 的用户
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var cmp func(p, q *TreeNode) bool
	cmp = func(p, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		if p.Val != q.Val {
			return false
		}
		if !cmp(p.Left, q.Right) {
			return false
		}
		if !cmp(p.Right, q.Left) {
			return false
		}
		return true
	}
	return cmp(root.Left, root.Right)
}
