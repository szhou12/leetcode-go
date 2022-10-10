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

func isPalindrome(head *ListNode) bool {
	// Edge Case: MUST have it o/w the algo is wrong when only one node
	if head.Next == nil {
		return true
	}

	// Step 1: Find the Middle Node
	// 同时要有一个指针记录slow前一个节点 (方便之后断链)
	var pre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		// reserve previous node before slow makes the move
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil

	// Step 2: cut the Linked List in half and reverse the second half
	cur1 := head
	cur2 := reverseList(slow)

	for cur1 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return true
}

// Iterative Way to reverse LL
// method: use pre, cur, next three pointers
func reverseList(head *ListNode) *ListNode {
	var next *ListNode
	cur := head
	var pre *ListNode

	for cur != nil {
		next = cur.Next
		cur.Next = pre // reverse connection
		pre = cur
		cur = next
	}

	return pre
}
