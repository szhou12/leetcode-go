package amz

import "container/heap"

/*
Given an array of positive integers, find max number of integers you can flip to negative such that the cumulative sum at each index is strictly positive.

Example:
nums = [5, 3, 1, 2]
flip =[5, -3, -1, 2]
cumsum = [5, 2, 1, 3]
Output: 2

Example:
nums = [6, 1, 1, 1, 4]
flip = [6, -1, -1, -1, -4]
cumsum= [6, 5, 4, 3, -1]
We want to recover -4 to 4
Output: 3
*/

/*
Idea: Greedy + Min Heap (Dynamically stores flipped numbers)
1. Greedily flip the current number to negative and keep track of cumsum.
2. If cumsum <= 0, pop the smallest negative number from min heap and recover it back to positive.
3. We keep track of # of flips along the way.
*/

func maxFlipKeepPosCumsum(nums []int) int {
	n := len(nums)

	minHeap := &PQCumsum{}
	heap.Init(minHeap)

	cumsum := 0
	res := 0

	for i := 0; i < n; i++ {
		cumsum += -nums[i]
		heap.Push(minHeap, -nums[i])

		// recover the smallest negative number until cumsum > 0
		for cumsum <= 0 {
			smallest := heap.Pop(minHeap).(int)
			cumsum -= smallest * 2
		}

		res = max(res, minHeap.Len()) // No need to track. The final output can just be the length of minHeap
	}
	return res
}

type PQCumsum []int

func (pq PQCumsum) Len() int {
	return len(pq)
}

func (pq PQCumsum) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQCumsum) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQCumsum) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQCumsum) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
