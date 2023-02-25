package leetcode

import (
	"container/heap"
	"math"
)

// leftMin[i] := min sum of n-len subseq in nums[0:i] = pick n smallest elements in nums[0:i] by using max heap
// rightMax[i] := max sum of n-len subseq in nums[i:3n-1] = pick n largest elements in nums[i:3n-1] by using min heap
func minimumDifference(nums []int) int64 {
	n := len(nums) / 3

	leftMin := make([]int, len(nums))
	maxHeap := &MaxPQ{}
	heap.Init(maxHeap)
	sum := 0
	for i := 0; i < len(nums); i++ {
		heap.Push(maxHeap, nums[i])
		sum += nums[i]
		if (*maxHeap).Len() > n {
			temp := heap.Pop(maxHeap)
			sum -= temp.(int)
		}
		leftMin[i] = sum
	}

	rightMax := make([]int, len(nums))
	minHeap := &MinPQ{}
	heap.Init(minHeap)
	sum = 0
	for i := len(nums) - 1; i >= 0; i-- {
		heap.Push(minHeap, nums[i])
		sum += nums[i]
		if (*minHeap).Len() > n {
			temp := heap.Pop(minHeap)
			sum -= temp.(int)
		}
		rightMax[i] = sum
	}

	res := math.MaxInt
	for i := n - 1; i < 3*n-n; i++ {
		res = min(res, leftMin[i]-rightMax[i+1])
	}
	return int64(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MinPQ []int

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq *MinPQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MinPQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}

type MaxPQ []int

func (pq MaxPQ) Len() int { return len(pq) }

func (pq MaxPQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq MaxPQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *MaxPQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MaxPQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
