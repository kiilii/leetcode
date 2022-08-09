package main

// 前序遍历

// 执行用时：0 ms, 在所有 Go 提交中击败了 100% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 23.37% 的用户
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ret []int
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)

	return ret
}
