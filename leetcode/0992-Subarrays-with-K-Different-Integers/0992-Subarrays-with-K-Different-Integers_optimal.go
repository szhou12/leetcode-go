package leetcode

// Optimal Solution O(n)
func subarraysWithKDistinct_optimal(nums []int, k int) int {
	n := len(nums)
	res := 0

	// atMost(k) - atMost(k-1) = exactly(k)
	upperLeft := atMost(nums, k-1)
	upperRight := atMost(nums, k)

	for i := 0; i < n; i++ {
		res += upperRight[i] - upperLeft[i]
	}
	return res
}

func atMost(nums []int, k int) []int {
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
		if !(right < n) && count <= k { // right出界了, windown内仍然 <= k distinct numbers
			upper[left] = right
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