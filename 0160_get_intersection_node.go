package main

// 执行用时：24 ms, 在所有 Go 提交中击败了 96.43% 的用户
// 内存消耗：7 MB, 在所有 Go 提交中击败了 38.11% 的用户
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var x, y = headA, headB
	for x != y {
		if x == nil {
			x = headB
		} else {
			x = x.Next
		}

		if y == nil {
			y = headA
		} else {
			y = y.Next
		}
	}
	return x
}
