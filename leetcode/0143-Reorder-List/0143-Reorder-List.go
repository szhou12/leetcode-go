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

func reorderList(head *ListNode) {
	// Step 1: find middle node as break point, cut to have two LL - LL1, LL2 (fast two steps, slow one step)
	temp := findMidNode(head) // 找到的是LL1最后一个节点
	head2 := temp.Next
	temp.Next = nil

	// Step 2: reverse LL2 (recursion)
	newHead2 := reverseList(head2)

	// Step 3: one node in LL1, connect to one node in LL2
	concatLL(head, newHead2)
}

func findMidNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// LC 206
func reverseList(head *ListNode) *ListNode {
	// base cases
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 具体操作: 先画图会比较容易想
func concatLL(head1 *ListNode, head2 *ListNode) {

	for head1 != nil && head2 != nil {
		h1Next := head1.Next
		h2Next := head2.Next

		head1.Next = head2
		head2.Next = h1Next

		head1 = h1Next
		head2 = h2Next
	}
}
