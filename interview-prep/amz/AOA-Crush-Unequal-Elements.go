package amz

import "container/heap"

/*
Given an array.
For one operation, you can pick two unequal elements and remove them from the array.
Find the minimum elements left in the array.

Leetcode 2856
*/

// Key Observation: cancel out elements with highest frequencies first
func crushUnequalElements(nums []int) int {

	// hashmap to store frequency of each element
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// max heap to store frequencies
	maxheap := &PQCrush{}
	heap.Init(maxheap)

	for _, f := range freq {
		heap.Push(maxheap, f)
	}

	// pop top 2 freqs and cancel each other once at a time until maxheap has no more than 1 type of freq left
	// we have to cancel out once at a time because in the next round, it's not necessarily true the top 2 freqs are still the same
	for maxheap.Len() > 1 {
		n1 := heap.Pop(maxheap).(int) // Top 1 freq
		n2 := heap.Pop(maxheap).(int) // Top 2 freq

		// freqs are dynamic, thus do NOT n1 -= n2
		// cancel out once at a time
		n1--
		n2--

		if n1 > 0 {
			heap.Push(maxheap, n1)
		}
		if n2 > 0 {
			heap.Push(maxheap, n2)
		}
	}

	// empty means all elements are cancelled out
	if maxheap.Len() == 0 {
		return 0
	}
	// one element left, its freq are the min elements left
	return heap.Pop(maxheap).(int)

}

type PQCrush []int

func (pq PQCrush) Len() int {
	return len(pq)
}

func (pq PQCrush) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PQCrush) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQCrush) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQCrush) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
