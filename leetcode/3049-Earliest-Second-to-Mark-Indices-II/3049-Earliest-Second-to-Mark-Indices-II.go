package leetcode

import "container/heap"

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	m := len(changeIndices)

	// prepend a placeholder to convert 0-indexed to 1-indexed
	nums = append([]int{0}, nums...)
	changeIndices = append([]int{0}, changeIndices...)

	// Binary Search
	left, right := 1, m
	for left < right {
		mid := left + (right-left)/2
		if valid(mid, nums, changeIndices) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// post-processing: check if the result given by loop is legal
	if valid(left, nums, changeIndices) {
		return left
	}

	return -1

}

func valid(t int, nums []int, changeIndices []int) bool {
	n := len(nums) - 1

	// Early Stop: ensure at t we have enough time to mark
	if t < n {
		return false
	}

	// Make all non-first-appearance nums idx in changeIndices to -1
	changeIndicesCopy := make([]int, len(changeIndices))
	copy(changeIndicesCopy, changeIndices)

	// first position a nums idx appears in changeIndices [nums idx -> changeIndices idx]
	firstOccur := make([]int, n+1)
	for i := 1; i <= t; i++ {
		// find the first position of a positive num value
		if firstOccur[changeIndices[i]] == 0 && nums[changeIndices[i]] > 0 {
			firstOccur[changeIndices[i]] = i
		} else {
			changeIndicesCopy[i] = -1
		}
	}


	// store num values that can be executed by set 0 operation
	minHeap := &PQ{}
	heap.Init(minHeap)

	// Greedy: loop backwards
	for i := t; i >= 1; i-- {
		numsIdx := changeIndicesCopy[i]
		if numsIdx == -1 {
			continue
		}
		heap.Push(minHeap, nums[numsIdx])
		marks := (t-i+1) - minHeap.Len() // count # of times to conduct mark operation
		if marks < minHeap.Len() { // if not enough mark ops, need to guarantee mark op first
			heap.Pop(minHeap)
		}
	}

	totalSetZero := 0
	setZeroCount := 0
	for minHeap.Len() > 0 {
		totalSetZero += heap.Pop(minHeap).(int)
		setZeroCount++
	}

	numsTotal := 0
	for _, val := range nums {
		numsTotal += val
	}

	// count total # of -1 operations = total seconds t - total marks n - # of set 0 operations
	totalMinusOne := t - n - setZeroCount

	return totalSetZero + totalMinusOne >= numsTotal

}

type PQ []int

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
