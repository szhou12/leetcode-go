package leetcode

import "container/heap"

func findMaximums(nums []int) []int {
	n := len(nums)

	prevSmallerOrEqual, nextSmaller := findSmaller(nums)

	maxHeap := &PQ{}
	heap.Init(maxHeap)

	for i, num := range nums {
		heap.Push(maxHeap, []int{num, i})
	}

	res := make([]int, 0)

	for k := 1; k <= n; k++ {
		for maxHeap.Len() > 0 {
			top := maxHeap.Top().([]int)
			if (nextSmaller[top[1]] - 1) - prevSmallerOrEqual[top[1]] >= k {
				break
			}
			heap.Pop(maxHeap)
		}

		res = append(res, maxHeap.Top().([]int)[0])

	}

	return res

}

func findSmaller(nums []int) ([]int, []int) {
	n := len(nums)

	// nextSmaller
	stack := make([]int, 0)

	nextSmaller := make([]int, n)

	for i := 0; i < n; i++ {
		nextSmaller[i] = n
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// prevSmallerOrEqual
	stack = make([]int, 0)

	prevSmallerOrEqual := make([]int, n)

	for i := 0; i < n; i++ {
		prevSmallerOrEqual[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			prevSmallerOrEqual[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

	}

	return prevSmallerOrEqual, nextSmaller
}

type PQ [][]int // [[value, index]]

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
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