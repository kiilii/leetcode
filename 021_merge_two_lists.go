package main

import "fmt"

func Test_021_mergeTwoLists() {
	l1 := genListNode([]int{1, 2, 4})
	l2 := genListNode([]int{1, 3, 4})
	// var l1 = []int{1, 2, 4}
	// var l2 = []int{1, 3, 4}

	fmt.Println(mergeTwoLists(l1, l2))
	// l1 = genListNode([]int{5})
	// l2 = genListNode([]int{1, 2, 4})
	// fmt.Println(mergeTwoLists2(l1, l2))

}

// func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val < list2.Val {
// 		return mergeTwoLists2(list1)
// 	}

// 	return list
// }

// 2023.04.29
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了 13.52%的用户
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var main = new(ListNode)
	var tmp = main

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	for {
		if list1.Val > list2.Val {
			tmp.Val = list2.Val
			list2 = list2.Next
		} else if list1.Val == list2.Val {
			tmp.Val = list2.Val
			list2 = list2.Next
		} else {
			tmp.Val = list1.Val
			list1 = list1.Next
		}

		if list1 == nil {
			tmp.Next = list2
			break
		}
		if list2 == nil {
			tmp.Next = list1
			break
		}
		tmp.Next = new(ListNode)
		tmp = tmp.Next
	}

	return main
}
