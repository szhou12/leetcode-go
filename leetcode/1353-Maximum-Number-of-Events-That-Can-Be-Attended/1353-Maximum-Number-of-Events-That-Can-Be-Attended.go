package leetcode

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	n := len(events)

	// STEP 1: sort by start date
	sort.Slice(events, func(i, j int) bool {
		// start day相同的任务，优先处理“更紧急”的(finish day早的)任务
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	// STEP 2: Compatible algo
	minHeap := &PQ{}
	heap.Init(minHeap)
	i := 0 // i-th event
	res := 0

	// day的物理意义: 哪些任务可以安排在这一天执行
	for day := 1; day <= int(1e+5); day++ { // O(1e+5)
		// For any given day, unlock all tasks that can be executed on this day (start_time <= day)
		// PQ sort them by finish time so as to always prioritize 'urgent' tasks (need to finish soon)
		for i < n && events[i][0] <= day { // O(nlogn)
			heap.Push(minHeap, events[i][1])
			i++
		}
		// Disqualify top tasks who already past due date (finish_time < day)
		for (*minHeap).Len() > 0 && (*minHeap)[0] < day { // O(nlogn)
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
