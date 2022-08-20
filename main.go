package main

// func main() {
// 	test()
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeListNode(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}

	var head, next *ListNode

	for i := 0; i < len(nodes); i++ {
		if i == 0 {
			head = new(ListNode)
			head.Val = nodes[i]
			next = head
		} else {
			if next.Next == nil {
				next.Next = new(ListNode)
				next = next.Next
				next.Val = nodes[i]
			}
		}
	}
	return head
}
