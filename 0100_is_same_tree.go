package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 48.29% 的用户
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}
