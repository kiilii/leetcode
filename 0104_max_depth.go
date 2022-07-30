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
// 内存消耗：4 MB, 在所有 Go 提交中击败了 98.68% 的用户
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	x := maxDepth(root.Left)
	y := maxDepth(root.Right)
	return max(x, y) + 1
}

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：4.1 MB, 在所有 Go 提交中击败了 6.34% 的用户
func maxDepth2(root *TreeNode) int {
	var cnt func(root *TreeNode, n int) int
	cnt = func(root *TreeNode, n int) int {
		if root == nil {
			return n
		}
		n += 1
		x := cnt(root.Left, n)
		y := cnt(root.Right, n)
		if x > y {
			return x
		}
		return y
	}
	return cnt(root, 0)
}
