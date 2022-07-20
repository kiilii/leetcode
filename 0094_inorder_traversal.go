package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 2023.07.20
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 18.49% 的用户
func inorderTraversal(root *TreeNode) []int {
	var x = make([]int, 0)
	if root == nil {
		return x
	}

	if root.Left != nil {
		x = append(x, inorderTraversal(root.Left)...)
	}

	x = append(x, root.Val)

	if root.Right != nil {
		x = append(x, inorderTraversal(root.Right)...)
	}

	return x
}
