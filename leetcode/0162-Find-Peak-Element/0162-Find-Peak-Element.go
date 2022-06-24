package leetcode

// FOUR CASES:
// case 1: A[m-1] > A[m], A[m] < A[m+1]  左/右边一定有peak
// case 2: A[m-1] > A[m], A[m] > A[m+1]  左边一定有peak; 根据条件1, end = mid
// case 3: A[m-1] < A[m], A[m] > A[m+1]  找到了peak
// case 4: A[m-1] < A[m], A[m] < A[m+1]  右边一定有peak; 根据条件2, start = mid
func findPeakElement(nums []int) int {
	// corner cases
	if nums == nil || len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid-1] {
			end = mid
		} else if nums[mid] < nums[mid+1] {
			start = mid
		} else {
			return mid
		}
	}

	if nums[start] > nums[end] {
		return start
	}

	return end
}
