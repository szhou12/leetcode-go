package leetcode

// Sliding Window: 模版Fix
func maximumSubarraySum(nums []int, k int) int64 {
	window := make(map[int]int)
	left, right := 0, 0
	count := 0 // record # of distinct numbers in the window
	sum := 0   // record current sum of numbers in the window
	res := 0

	for right < len(nums) {
		rightElement := nums[right] // 吃
		right++
		sum += rightElement
		window[rightElement]++
		if window[rightElement] == 1 { // found a distinct number
			count++
		}

		for right-left >= k {
			// update result: only when we have k distinct numbers in the window
			if count == k {
				res = max(res, sum)
			}

			leftElement := nums[left] // 吐
			left++
			sum -= leftElement
			if window[leftElement] == 1 {
				count--
			}
			window[leftElement]--
		}
	}

	return int64(res)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
