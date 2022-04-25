package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l1 := genListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l1 := genListNode([]int{1, 0, 6, 9, 0, 1})
	l2 := genListNode([]int{5, 6, 4})

	// fmt.Println(num)
	addTwoNumbers(l1, l2)
}

func genListNode(nums []int) *ListNode {
	var l = new(ListNode)
	var tmp = l

	for i := 0; i < len(nums); i++ {
		tmp.Val = nums[i]
		if i < len(nums)-1 {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		}
	}

	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ln = new(ListNode)
	var tmp = ln

	for {
		if l1 != nil {
			tmp.Val = l1.Val + tmp.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val = l2.Val + tmp.Val
			l2 = l2.Next
		}

		if tmp.Val >= 10 {
			tmp.Val = tmp.Val % 10
			tmp.Next = &ListNode{Val: 1}
		} else {
			if l1 == nil && l2 == nil {
				break
			}
			tmp.Next = new(ListNode)
		}
		tmp = tmp.Next
	}
	return ln
}
