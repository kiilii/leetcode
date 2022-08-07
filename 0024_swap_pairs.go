package main

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了 58.93%的用户
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var x = &ListNode{Val: 0, Next: head}

	var h, e, o = x, x.Next, x.Next.Next

	for e != nil || o != nil {
		h.Next = o
		e.Next = o.Next
		o.Next = e
		// 移动两个节点
		h = e
		if e.Next == nil || e.Next.Next == nil {
			break
		}
		e = e.Next
		o = e.Next
	}

	return x.Next
}
