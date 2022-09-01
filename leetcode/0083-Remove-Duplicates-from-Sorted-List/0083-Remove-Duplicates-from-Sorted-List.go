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

func deleteDuplicates(head *ListNode) *ListNode {
	// Edge Cases
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	// 如果终止条件是 cur.Next == nil, 写完要检查两种情况：
	// 1. 最后两个元素相同
	// 2. 最后两个元素不同
	for cur.Next != nil { // avoid NPE
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next //断开与下一个的链接，与下下一个链接
		} else {
			cur = cur.Next
		}
	}
	return head
}

// WRONG SOLUTION
// Input: [1,1,2,3,3,3]
// Output: [1,2,3,3,3]
func deleteDuplicates_WRONG(head *ListNode) *ListNode {
	// Edge cases
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: -1, Next: nil}
	cur := dummy
	for head != nil {
		if head.Val != cur.Val {
			cur.Next = head // 链接每个重复中的第一个元素,后面重复的没有断开
			cur = cur.Next
		}
		head = head.Next

	}

	return dummy.Next

}
