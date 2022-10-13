package leetcode

import "container/heap"

type PQ [][2]int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][1] < pq[j][1]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.([2]int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: use a map to record frequency
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// Step 2: min heap
	minHeap := &PQ{}
	heap.Init(minHeap)

	for key, val := range freqMap {
		heap.Push(minHeap, [2]int{key, val})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(minHeap).([2]int)[0]
	}
	return res
}
