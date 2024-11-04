package leetcode

import "container/heap"

// Solution: Priority Queue - min heap
func findKthLargest_PQ(nums []int, k int) int {

	minHeap := &PQ{}
	heap.Init(minHeap)

	for _, num := range nums {
		if (*minHeap).Len() < k {
			heap.Push(minHeap, num)
		} else {
			if num > (*minHeap)[0] {
				heap.Pop(minHeap)
				heap.Push(minHeap, num)
			}
		}
	}

	res := heap.Pop(minHeap).(int)

	return res
}

type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
