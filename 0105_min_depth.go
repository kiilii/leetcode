package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 执行用时：164 ms, 在所有 Go 提交中击败了 56.93% 的用户
// 内存消耗：17.9 MB, 在所有 Go 提交中击败了 77.79% 的用户
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	var min = func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
