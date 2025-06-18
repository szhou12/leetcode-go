package main

import (
	"container/heap"
	"strings"
)

/*
*
Max heap to store each letter and its frequency.
Always pop the top 2 highest freq letters, concat them
Until there is < 2 letters left in the max heap
if the left-over letter still has freq > 1, we are unable to reogranize the string
o/w, we return the result
*/
func reorganizeString(s string) string {
	// use a map to count freq
	freq := make(map[int]int)
	for _, char := range s {
		charInt := int(char - 'a')
		freq[charInt]++
	}

	maxHeap := &PQ{}
	heap.Init(maxHeap)

	for charInt, count := range freq {
		heap.Push(maxHeap, []int{charInt, count})
	}

	var sb strings.Builder

	for maxHeap.Len() > 1 {
		top1 := heap.Pop(maxHeap).([]int)
		top2 := heap.Pop(maxHeap).([]int)

		sb.WriteByte(byte(top1[0] + 'a'))
		sb.WriteByte(byte(top2[0] + 'a'))

		top1[1]--
		top2[1]--
		
		if top1[1] > 0 {
			heap.Push(maxHeap, top1)
		}

		if top2[1] > 0 {
			heap.Push(maxHeap, top2)
		}

	}

	if maxHeap.Len() > 0 {
		top := heap.Pop(maxHeap).([]int)
		if top[1] > 1 {
			return ""
		} else {
			sb.WriteByte(byte(top[0] + 'a'))
		}
	}

	return sb.String()

}

type PQ [][]int // [[charInt, freq], ...]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][1] > pq[j][1]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}

func (pq PQ) Top() interface{} {
	return pq[0]
}
