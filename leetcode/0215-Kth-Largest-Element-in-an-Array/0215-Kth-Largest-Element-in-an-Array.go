package leetcode

import (
	"container/heap"
	"math"
)

// Solution: Binary Searcy by Value
func findKthLargest_BS(nums []int, k int) int {
	left := math.MinInt / 2
	right := math.MaxInt / 2

	for left < right { // stop at left == right
		mid := right - (right-left)/2
		if count(nums, mid) >= k {
			left = mid // this can lead to dead loop when left=0, right=1, then mid = left + (right-left)/2 = 0. Hence, mid = right - (right-left)/2
		} else {
			right = mid - 1
		}
	}
	// return either left or right bc while loop stops at left == right
	return left
}

func count(nums []int, t int) int {
	// count the # of elements in nums that are >= t
	res := 0
	for _, num := range nums {
		if num >= t {
			res++
		}
	}

	return res
}

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
