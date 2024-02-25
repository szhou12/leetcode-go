package leetcode

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	n := len(events)

	// STEP 1: sort by start date
	sort.Slice(events, func(i, j int) bool {
		// No need to handle events[i][0] == events[j][0] bc PQ will handle it
		return events[i][0] < events[j][0]
	})

	// STEP 2: Compatible algo
	minHeap := &PQ{}
	heap.Init(minHeap)
	i := 0 // i-th event
	res := 0

	for deadline := 1; deadline <= int(1e+5); deadline++ {
		// For any given deadline day, select all tasks that can start before the deadline as candidates
		// PQ sort them by finish time so as to always prioritize 'urgent' tasks (need to finish soon)
		for i < n && events[i][0] <= deadline {
			heap.Push(minHeap, events[i][1])
			i++
		}
		// Disqualify top tasks who already past due date (finish_time < deadline)
		for (*minHeap).Len() > 0 && (*minHeap)[0] < deadline {
			heap.Pop(minHeap)
		}
		if (*minHeap).Len() > 0 {
			heap.Pop(minHeap)
			res++
		}
	}
	return res

}

type PQ []int // [start_date1, start_date2, ...]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
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
