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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// Step 1: 分别算出 两条链的长度 (nodes总数)
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	for curA != nil {
		lenA++
		curA = curA.Next
	}
	for curB != nil {
		lenB++
		curB = curB.Next
	}

	// Step 2: 算出 长度差, 较长的链上的指针 先走 长度差 个nodes
	var stepsAhead int
	var fast, slow *ListNode
	if lenA > lenB {
		stepsAhead = lenA - lenB
		fast, slow = headA, headB
	} else {
		stepsAhead = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < stepsAhead; i++ { // fast指针提前走
		fast = fast.Next
	}

	// Step 3: fast, slow同时走，一次一步，相遇时非空节点则是所求，否则相遇时为nil，则没有交叉节点
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
