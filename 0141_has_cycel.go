package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 执行用时：8 ms, 在所有 Go 提交中击败了 34.65% 的用户
// 内存消耗：4.2 MB, 在所有 Go 提交中击败了 64.91% 的用户
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow, fast = head, head.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
