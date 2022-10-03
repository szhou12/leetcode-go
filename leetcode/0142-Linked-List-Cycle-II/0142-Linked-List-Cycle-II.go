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

func detectCycle(head *ListNode) *ListNode {

	// Step 1: check whether has a cycle; and if so, find slow's current position
	cyclic, slow := hasCycle(head)
	if !cyclic {
		return nil
	}

	// Step 2: both fast, slow walk at pace = 1. WHne they meet, the meeting node is the start of cycle
	fast := head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func hasCycle(head *ListNode) (bool, *ListNode) {
	// Edge Case
	if head == nil {
		return false, nil
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true, slow
		}
	}
	return false, nil
}
