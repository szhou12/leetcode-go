package sorting

import (
	"math/rand"
	"sync"
)

func QuickSort(nums []int, isParallel bool) {
	if len(nums) <= 1 {
		return
	}

	if !isParallel {
		quickSortSequential(nums, 0, len(nums)-1)
	} else {
		quickSortParallel(nums, 0, len(nums)-1)
	}

}



func partition(nums []int, left int, right int) int {
	// choose a pivot
	pivotIndex := getPivotIndexLast(left, right)
	pivot := nums[pivotIndex]
	// swap pivot to the end
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]

	// two pointers: partition respect to pivot
	l, r := left, right-1
	for l <= r {
		if nums[l] <= pivot {
			l++
		} else if nums[r] > pivot {
			r--
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	// put pivot to l bc now l points to the first element > pivot
	nums[l], nums[right] = nums[right], nums[l]
	return l

}

func getPivotIndexLast(left int, right int) int {
	return right
}

func getPivotIndexRandom(left int, right int) int {
	return rand.Intn(right-left+1) + left
}

/*
* Sequential Quick Sort: not leveraging parallelism
*/
func quickSortSequential(nums []int, left int, right int) {
	// base case
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)
	quickSortSequential(nums, left, pivot-1)
	quickSortSequential(nums, pivot+1, right)
}


const (
	// THRESHOLD - arrays smaller than this are quick sorted sequentially
	THRESHOLD = 10
)

/*
* Parallel Quick Sort: leveraging parallelism by using goroutines and waitgroup
*/
func quickSortParallel(nums []int, left int, right int) {
	if left >= right {
		return
	}

	n := right - left + 1
	if n <= THRESHOLD {
		quickSortSequential(nums, left, right)
		return
	}

	pivot := partition(nums, left, right)

	// Use WaitGroup to implement fork/join pattern
	var wg sync.WaitGroup

	// Fork: spwan a new thread (goroutine) for the left partition
	if pivot > left {
		wg.Add(1)
		go func() {
			defer wg.Done()
			quickSortParallel(nums, left, pivot-1)
		}()
	}

	// current thread continues with the right partition
	if pivot < right {
		quickSortParallel(nums, pivot+1, right)
	}

	// Join: wait for the spawned goroutine to finish
	wg.Wait()

}
