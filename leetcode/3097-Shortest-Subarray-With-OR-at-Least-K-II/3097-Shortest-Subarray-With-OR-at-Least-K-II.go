package leetcode

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if valid(nums, mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// result may not exist
	if valid(nums, left, k) {
		return left
	}
	return -1
}

// Sliding Window + bitwise OR sum trick using count[] of len=31
func valid(nums []int, l int, k int) bool {
	n := len(nums)
	count := make([]int, 31)
	// preparation: update count[] for first l-2 elements [0:l-2] (inclusive)
	for i := 0; i < l-1; i++ {
		for b := 0; b < 31; b++ {
			count[b] += (nums[i] >> b) & 1
		}
	}

	// sliding window
	// move right pointer
	for right := l - 1; right < n; right++ {
		// 吃
		for b := 0; b < 31; b++ {
			count[b] += (nums[right] >> b) & 1
		}

		// update: bitwise OR sum in the sliding window
		sum := 0
		for b := 0; b < 31; b++ {
			if count[b] > 0 {
				sum += (1 << b)
			}
		}
		if sum >= k {
			return true
		}

		// 吐
		for b := 0; b < 31; b++ {
			count[b] -= (nums[right-l+1] >> b) & 1
		}
	}

	return false
}
