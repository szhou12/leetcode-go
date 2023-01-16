package leetcode

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m := len(nums1)
	n := len(nums2)

	left := nums1[0] + nums2[0]
	right := nums1[m-1] + nums2[n-1]

	for left < right {
		mid := left + (right-left)/2
		count := countLessOrEqual(nums1, nums2, mid)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	kSmallestSum := left

	// NOTE: separate smaller sums and equal sums bc otherwise, adding equal sums may exceeed k elements that we want to return
	// when there are too many equal sums
	smaller := make([][]int, 0)
	equal := make([][]int, 0)
	for i := 0; i < len(nums1); i++ { // O(n)
		j := 0
		for j < len(nums2) && nums1[i]+nums2[j] <= kSmallestSum { // O(k)
			if nums1[i]+nums2[j] < kSmallestSum {
				smaller = append(smaller, []int{nums1[i], nums2[j]})
			} else {
				equal = append(equal, []int{nums1[i], nums2[j]})
			}
			j++
		}
	}

	// NOTE: equal array may not have (k - len(smaller)) this many elements
	residual := min(k-len(smaller), len(equal))
	for i := 0; i < residual; i++ {
		smaller = append(smaller, equal[i])
	}

	return smaller
}

func countLessOrEqual(nums1 []int, nums2 []int, mid int) int {
	m, n := len(nums1), len(nums2)
	i, j := m-1, 0

	count := 0
	for i >= 0 && j < n {
		if nums1[i]+nums2[j] <= mid {
			count += (i + 1)
			j++

		} else {
			i--
		}

	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// naive solution: Max Heap to find k smallest pairs
// Time Limit Exceed!!!
func kSmallestPairs_PQ(nums1 []int, nums2 []int, k int) [][]int {
	maxHeap := &PQ{}
	heap.Init(maxHeap)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			pair := []int{nums1[i], nums2[j]}
			heap.Push(maxHeap, pair)
			if maxHeap.Len() > k {
				heap.Pop(maxHeap)
			}
		}
	}

	// res := make([][]int, k)
	// for i := k-1; i >= 0; i-- {
	//     res[i] = heap.Pop(maxHeap).([]int)
	// }

	res := make([][]int, 0)
	for maxHeap.Len() > 0 {
		res = append([][]int{heap.Pop(maxHeap).([]int)}, res...)
	}

	return res

}

type PQ [][]int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i][0]+pq[i][1] > pq[j][0]+pq[j][1]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	(*pq) = append(*pq, x.([]int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
