package leetcode

import "github.com/szhou12/leetcode-go/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structures.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// base cases
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// recursion
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
