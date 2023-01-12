package leetcode

import "container/heap"

// Optimal Solution: Binary Search
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]
	right := matrix[n-1][n-1]

	for left < right {
		mid := left + (right-left)/2
		count := countLessOrEqual(matrix, mid) // # of matrix elements <= mid
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// O(m+n) bc it traverses diagonally
func countLessOrEqual(matrix [][]int, mid int) int {
	n := len(matrix)
	// Starting pos: bottom-left corner / top-right corner
	i, j := n-1, 0

	res := 0

	// Traverses diagonally: i一直减小, j一直增加, 直至出界
	for i >= 0 && j < n {
		if matrix[i][j] <= mid { // 说明 (i,j) 上面包括 (i, j)自己 的值都 <= mid, 所以都算入res, 且向右走得到大一点的值
			res += (i + 1)
			j++
		} else { // (i, j) 过于大, 不合要求, 向上走得到小一点的值
			i--
		}
	}
	return res
}

// Priority Queue: max Heap
func kthSmallest_PQ(matrix [][]int, k int) int {
	maxHeap := &PQ{}
	heap.Init(maxHeap)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if (*maxHeap).Len() < k {
				heap.Push(maxHeap, matrix[i][j])
			} else if matrix[i][j] < (*maxHeap)[0] {
				heap.Pop(maxHeap)
				heap.Push(maxHeap, matrix[i][j])
			} else {
				break
			}
		}
	}
	return heap.Pop(maxHeap).(int)
}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
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
