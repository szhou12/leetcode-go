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

// My Solution
func removeElements(head *ListNode, val int) *ListNode {
	// Edge case
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: nil}
	cur := head
	prev := dummy

	for cur.Next != nil { // STOP at tail node. ie. Tail node NOT processed in the loop
		if cur.Val != val { // ONLY connect cur node to linked list when cur node is not the target
			prev.Next = cur
			prev = prev.Next
		}
		cur = cur.Next
	}

	// post-processing
	if cur.Val != val {
		prev.Next = cur
	} else {
		prev.Next = nil
	}

	return dummy.Next
}

// Optimal Solution - faster than My Solution
// BUT I find it harder to understand
func removeElements_optimal(head *ListNode, val int) *ListNode {
	// Edge case
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: head} // HERE: dummy.Next connects to head
	cur := head
	prev := dummy
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
