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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head} // 虚拟头结点, 防止因为删除的是第一个节点导致的空指针
	x := findFromEnd(dummy, n+1)            // 删除倒数第 n 个，要先找倒数第 n + 1 个节点
	x.Next = x.Next.Next
	return dummy.Next
}

// 返回倒数第k个node
func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	// p1 先走 k 步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := head
	// p1 和 p2 同时走 n - k 步, 即, p1走到底
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// p2 现在指向第 n - k 个节点, 即, 倒数第k个节点
	return p2
}

// k = 2
// d -> 1 -> 2 -> 3 -> 4 -> 5 ->
// p1
// p2             p1
//                p2             p1
