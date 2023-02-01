package leetcode

func findKthPositive(arr []int, k int) int {
	left := 1
	right := arr[len(arr)-1] + k // kth missing is at most k elements after last element in arr (when arr consecutive)

	for left < right {
		mid := right - (right-left)/2

		// bc we're guessing whether mid == kth missing, here we want to know wether there're k-1 missing # < mid
		totals := mid - 1                // total numbers in [1, mid-1]
		presents := lowerBound(arr, mid) // 利用 arr 单增的性质 找到 arr中第一个 >= mid的元素index j, 则存在于arr中 < mid的元素有 j-0=j个
		missing := totals - presents

		if missing <= k-1 {
			left = mid
		} else { // mising > k-1 说明 已经至少有 k missing numbers < mid, mid偏大, 一定不是答案
			right = mid - 1
		}
	}

	return left
}

// find 1st index whose element >= target
func lowerBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// post-processing
	if nums[left] >= target {
		return left
	} else {
		return left + 1
	}
}
