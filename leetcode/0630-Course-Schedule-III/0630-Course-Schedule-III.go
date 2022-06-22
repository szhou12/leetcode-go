package leetcode

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	// sort by finish time in increasing order
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	// total time
	time := 0
	maxHeap := &PQ{}
	heap.Init(maxHeap)
	for _, course := range courses {
		if time+course[0] <= course[1] {
			heap.Push(maxHeap, course[0])
			time += course[0]
		} else if (*maxHeap).Len() > 0 && (*maxHeap)[0] > course[0] {
			time -= heap.Pop(maxHeap).(int) - course[0]
			heap.Push(maxHeap, course[0])
		}
	}

	return (*maxHeap).Len()

}

type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp

}
