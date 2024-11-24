package leetcode

import (
	"container/heap"
	"sort"
)

func kSum(nums []int, k int) int64 {
	maxSum := int64(0)
	for _, x := range nums {
		if x > 0 {
			maxSum += int64(x)
		}
	}
	//edge case
	if k == 1 {
		return maxSum
	}

	// convert nums[] to absolute values of nums[]
	for i, x := range nums {
		if x < 0 {
			nums[i] = -x
		}
	}

	// sort nums[] in ascending order
	sort.Ints(nums)

	minHeap := &PQ{}
	heap.Init(minHeap)

	heap.Push(minHeap, []int{nums[0], 0})

	for t := 0; t < k-1; t++ {
		cur := heap.Pop(minHeap).([]int)
		sum, i := cur[0], cur[1]

		if t == k-2 {
			return maxSum - int64(sum)
		}

		if i+1 < len(nums) {
			heap.Push(minHeap, []int{sum + nums[i+1], i + 1})
			heap.Push(minHeap, []int{sum - nums[i] + nums[i+1], i + 1})
		}

	}

	// wont reach here
	return -1
}

type PQ [][]int // [[sequence sum ending at i], i]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
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
