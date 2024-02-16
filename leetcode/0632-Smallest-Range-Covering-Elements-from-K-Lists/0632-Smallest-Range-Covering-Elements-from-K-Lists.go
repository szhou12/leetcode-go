package leetcode

import (
	"container/heap"
	"math"
)

func smallestRange(nums [][]int) []int {
	n := len(nums)

	minHeap := &PQ{}
	heap.Init(minHeap)
	maxVal := math.MinInt

	for i, list := range nums {
		heap.Push(minHeap, []int{list[0], i})
		maxVal = max(maxVal, list[0])
	}

	// indices pointing to elements that are currently in the heap
	pointers := make([]int, n)

	rangeSize := math.MaxInt
	var res []int
	for {
		smallest := heap.Pop(minHeap).([]int)
		// update: range, res
		curRange := maxVal - smallest[0]
		if curRange < rangeSize {
			rangeSize = curRange
			res = []int{smallest[0], maxVal}
		}
		
		// check if any list is exhausted
		i := smallest[1]
		pointers[i]++
		if pointers[i] == len(nums[i]) {
			break
		}

		// push new element to heap
		heap.Push(minHeap, []int{nums[i][pointers[i]], i})
		maxVal = max(maxVal, nums[i][pointers[i]])
	}

	return res
	
}

type PQ [][]int // [[value, list index]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	res := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return res
}

func (pq *PQ) Peek() interface{} {
	return (*pq)[0]
}
