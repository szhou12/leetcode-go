package leetcode

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	minHeap := &PQ{}
	heap.Init(minHeap)

	for _, node := range lists {
		if node != nil {
			heap.Push(minHeap, node)
		}
	}

	for minHeap.Len() > 0 {
		// pop the currently smallest value node
		node := heap.Pop(minHeap).(*ListNode)
		cur.Next = node

		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}

		cur = cur.Next
	}

	return dummy.Next
}

type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
