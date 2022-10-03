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

// Recursion
// Return: new head which is the last node of input LL
func reverseList(head *ListNode) *ListNode {
	// Base Cases: 注意！这里有两个Base Cases
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head // node1 -> node2 => node2 -> node1
	head.Next = nil       // 断开链接
	return newHead

}
