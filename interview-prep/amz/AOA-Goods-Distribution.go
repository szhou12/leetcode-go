package amz

import (
	"container/heap"
	"sort"
)

/*
Given an array of size n,
Minimize the difference between the max and min elements
We can achieve this by this operation: pick two elements, decrement one by 1 and increment the other by 1
Find the minimum # of operations to meet the requirement

Example:
Input: [1, 2, 5, 8]
pick [1,8]: [1+1, 2, 5, 8-1] => [2, 2, 5, 7]
pick [2,7]: [2+1, 2, 5, 7-1] => [3, 2, 5, 6]
pick [2,6]: [3, 2+1, 5, 6-1] => [3, 3, 5, 5]
pick [3,5]: [3, 3+1, 5-1, 5] => [3, 4, 4, 5]
pick [3,5]: [3+1, 4, 4, 5-1] => [4, 4, 4, 4]
Output: 5

*/

// Solution 1: MaxHeap, MinHeap
// as long as (MaxHeap.Peek() - MinHeap.Peek() > 1), we can max-1 and put back to MaxHeap, and min+1 and put back to MinHeap, count++
func minOpsToMinimizeDifference_pq(nums []int) int {
	minHeap := &PQ1{}
	heap.Init(minHeap)
	maxHeap := &PQ2{}
	heap.Init(maxHeap)

	for _, num := range nums {
		heap.Push(minHeap, num)
		heap.Push(maxHeap, num)
	}

	count := 0
	for (*maxHeap)[0]-(*minHeap)[0] > 1 {
		count++

		maximum := heap.Pop(maxHeap).(int)
		minimum := heap.Pop(minHeap).(int)

		heap.Push(maxHeap, maximum-1)
		heap.Push(minHeap, minimum+1)
	}

	return count

}

// Solution 2: Sorting + Two Pointers
// Key Observation: At the end of the day, the resulting array will have either diff = 0 or diff = 1
// NOTE: this solution is NOT correct
func minOpsToMinimizeDifference(nums []int) int {
	n := len(nums)

	sort.Ints(nums)

	left := 0      // points to the globally smallest element
	right := n - 1 // points to the globally largest element
	count := 0
	for left < right && nums[right]-nums[left] > 1 {
		count++

		nums[left]++
		nums[right]--

		// left always points to the smallest element
		if nums[left] > nums[left+1] {
			left++
		}

		// right always points to the largest element
		if nums[right] < nums[right-1] {
			right--
		}

	}

	return count
}

type PQ1 []int

func (pq PQ1) Len() int { return len(pq) }
func (pq PQ1) Less(i, j int) bool {
	return pq[i] < pq[j] // min heap
}
func (pq PQ1) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ1) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ1) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}

type PQ2 []int

func (pq PQ2) Len() int { return len(pq) }
func (pq PQ2) Less(i, j int) bool {
	return pq[i] > pq[j] // max heap
}
func (pq PQ2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ2) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ2) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
