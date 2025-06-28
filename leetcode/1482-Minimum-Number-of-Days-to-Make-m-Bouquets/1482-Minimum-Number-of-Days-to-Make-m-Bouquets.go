package leetcode

func minDays(bloomDay []int, m int, k int) int {
	left, right := bloomDay[0], bloomDay[0]
	for _, day := range bloomDay {
		left = min(left, day) // at least need to wait the earliest flower to bloom
		right = max(right, day) // all flowers bloom on this day
	}

	for left < right {
		mid := left + (right - left) / 2
		if isOk(bloomDay, m, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// post-processing: not necessarily have a valid answer
	if isOk(bloomDay, m, k, left) {
		return left
	}

	return -1
}

// Check if we can make >= m bonquets in targetDays 
func isOk(bloomDay []int, m int, k int, targetDays int) bool {
	n := len(bloomDay)
	count := 0 // number of flowers in current bonquet
	bonquets := 0

	for i := 0; i < n; i++ {
		if bloomDay[i] <= targetDays {
			count++
		} else {
			count = 0 // cause we need to have consecutive flowers, once break, reset count
		}

		if count == k {
			bonquets++
			count = 0
		}
	}

	return bonquets >= m
}