package leetcode

// 						idx 		val
// 					-----------     -------------
// nums 			     n          decrement cnts
// changeIndices         m          nums idx
// lastOccur          nums idx      changeIndices idx
func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	m := len(changeIndices)

	// prepend a placeholder to convert 0-indexed to 1-indexed
	nums = append([]int{0}, nums...)
	changeIndices = append([]int{0}, changeIndices...)

	// Binary Search
	left, right := 1, m
	for left < right {
		mid := left + (right-left) / 2
		if valid(mid, nums, changeIndices) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// check if the result given by loop is legal
	if valid(left, nums, changeIndices) {
		return left
	}
	return -1
}

// Greedy: check if at time t, all nums val can be reduced to 0 and all nums idx can be marked
func valid(t int, nums []int, changeIndices []int) bool {
	n := len(nums) - 1

	// last position a nums idx appears in changeIndices [nums idx -> changeIndices idx]
	lastOccur := make([]int, n+1)
	// t in [1, m]
	for i := 1; i <= t; i++ {
		lastOccur[changeIndices[i]] = i
	}

	// Early Stop: 
	// if not all nums idx are captured in lastOccur, meaning we can't mark all indices at time t
	for i := 1; i <= n; i++ {
		if lastOccur[i] == 0 {
			return false
		}
	}

	decrementCount := 0
	// t in [1, m]
	for i := 1; i <= t; i++ {
		numsIdx := changeIndices[i]
		// if i not the last position a num appears in changeIndices, accumulate decrement count+1
		if i != lastOccur[numsIdx] {
			decrementCount++
		} else {
			decrementCount -= nums[numsIdx]
			if decrementCount < 0 {
				return false
			}
		}
	}

	return true

}