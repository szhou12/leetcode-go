package leetcode

// My Solution
// Find the first index (shortest sliding window) that contains exactly k distinct numbers
// FInd the last index (longest sliding window) that contains exactly k distinct numbers
// The difference between the two = # of subarrays given a left index
func subarraysWithKDistinct(nums []int, k int) int {
	n := len(nums)
	res := 0
	lower := lowerBound(nums, k)
	upper := upperBound(nums, k)

	for i := 0; i < n; i++ {
		if lower[i] == -1 {
			continue
		}
		res += upper[i] - lower[i]
	}

	return res
}

// lower bound:
// for each index i as left, find the first index as right that makes subarray of exactly k distinct numbers
func lowerBound(nums []int, k int) []int {
	n := len(nums)
	lower := make([]int, n)
	window := make(map[int]int)
	count := 0
	right := 0

	for left := 0; left < n; left++ {
		// 吃
		for right < n && count < k {
			rightElement := nums[right]
			if window[rightElement] == 0 {
				count++
			}
			window[rightElement]++
			right++
		}

		// update lower bound
		if !(right < n) && count < k {
			lower[left] = -1
		} else if (right < n) && !(count < k) {
			lower[left] = right - 1
		} else {
			lower[left] = right - 1
		}

		// 吐
		leftElement := nums[left]
		window[leftElement]--
		if window[leftElement] == 0 {
			count--
		}
	}

	return lower
}

// upper bound
// for each index i as left, find the last index as right that makes subarray of exactly k distinct numbers
// NOTE: each resulting right = the first index that makes subarray of > k distinct numbers
func upperBound(nums []int, k int) []int {
	n := len(nums)
	upper := make([]int, n)
	window := make(map[int]int) // record appearances of each distinct number within the window
	count := 0 // count distinct numbers within the window
	right := 0

	for left := 0; left < n; left++ {
		// 吃
		for right < n && count <= k {
			rightElement := nums[right]
			if window[rightElement] == 0 {
				count++
			}
			window[rightElement]++
			right++
		}

		// update upper bound
		if !(right < n) && count <= k {
			if count < k {
				upper[left] = -1 // right出界了都没凑齐k个distinct number
			} else { // count == k
				upper[left] = right // right出界 和 凑齐k个 同时发生
			}
		} else if (right < n) && !(count <= k) {
			upper[left] = right - 1 // right-1是因为在超过k个之后right又多走了一步,right-1才是第一个超过k个的index
		} else { // !(right < n) && !(count <= k)
			upper[left] = right - 1 // right-1是因为在超过k个之后right又多走了一步
		}

		// 吐
		leftElement := nums[left]
		window[leftElement]--
		if window[leftElement] == 0 {
			count--
		}
	}

	return upper
}
