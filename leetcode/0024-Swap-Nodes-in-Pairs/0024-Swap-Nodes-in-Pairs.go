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

func swapPairs(head *ListNode) *ListNode {
	// Base case
	if head == nil || head.Next == nil {
		return head
	}

	// Recursion
	first := head
	second := head.Next
	others := head.Next.Next
	second.Next = first
	first.Next = swapPairs(others)
	return second
}
