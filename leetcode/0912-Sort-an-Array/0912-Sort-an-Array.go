package leetcode

import "math/rand"

func sortArray(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}

	// Merge Sort
	// helper := make([]int, len(nums))
	// mergeSort(nums, helper, 0, len(nums)-1)

	// Quick Sort
	quickSort(nums, 0, len(nums)-1)

	return nums
}

/*
* Merge Sort implementation
 */
func mergeSort(array []int, helper []int, left int, right int) {
	// base case
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(array, helper, left, mid)
	mergeSort(array, helper, mid+1, right)
	merge(array, helper, left, mid, right)
}

func merge(array []int, helper []int, left int, mid int, right int) {
	// copy to helepr
	for i := left; i <= right; i++ {
		(helper)[i] = (array)[i]
	}

	l := left
	r := mid + 1
	for l <= mid && r <= right {
		if (helper)[l] <= (helper)[r] {
			(array)[left] = (helper)[l]
			left += 1
			l += 1
		} else {
			(array)[left] = (helper)[r]
			left += 1
			r += 1
		}
	}

	// post-processing
	for l <= mid {
		(array)[left] = (helper)[l]
		left += 1
		l += 1
	}

}

/*
* Quick Sort implementation
 */
func quickSort(array []int, left int, right int) {
	// base case
	if left >= right {
		return
	}

	pivot := partition(array, left, right)
	quickSort(array, left, pivot-1)
	quickSort(array, pivot+1, right)
}

func partition(array []int, left int, right int) int {
	pIndex := getPivotIndex(left, right)
	pivot := array[pIndex]
	array[right], array[pIndex] = array[pIndex], array[right]

	l := left
	r := right - 1
	for l <= r {
		if array[l] <= pivot {
			l += 1
		} else if array[r] > pivot {
			r -= 1
		} else {
			array[l], array[r] = array[r], array[l]
			l += 1
			r -= 1
		}
	}

	array[l], array[right] = array[right], array[l]
	return l
}

func getPivotIndex(left int, right int) int {
	return left + rand.Intn(right-left+1)
}
