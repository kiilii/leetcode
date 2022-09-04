package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 执行用时：4 ms, 在所有 Go 提交中击败了 97.47% 的用户
// 内存消耗：4.5 MB, 在所有 Go 提交中击败了 63.93% 的用户
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for head != nil && head.Val == val {
		head = head.Next
	}
	var next = head

	for next != nil && next.Next != nil {
		if next.Next.Val == val {
			next.Next = next.Next.Next
		} else {
			next = next.Next
		}
	}
	return head
}
