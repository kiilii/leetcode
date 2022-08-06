package main

// 2023.07.21
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3 MB, 在所有 Go 提交中击败了 59.81% 的用户
func deleteDuplicates(head *ListNode) *ListNode {
	var fast, slow = head, head

	for slow != nil {
		if fast.Next == nil {
			slow.Next = fast.Next
			break
		} else {
			fast = fast.Next
			if slow.Val == fast.Val {
				continue
			}
			slow.Next = fast
			slow = slow.Next
		}
	}

	return head
}
