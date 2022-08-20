package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了 5.69% 的用户
// 递归法
func reverseListV1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var x func(node, next *ListNode) (*ListNode, *ListNode)
	x = func(node, next *ListNode) (*ListNode, *ListNode) {
		if next == nil {
			return node, node
		}
		head, tail := x(next, next.Next)
		tail.Next = node
		node.Next = nil

		return head, node
	}

	node, _ := x(head, head.Next)
	return node
}
