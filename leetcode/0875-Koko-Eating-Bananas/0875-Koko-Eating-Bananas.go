package leetcode

import "math"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, math.MaxInt // NOTE: right can be limited to max(piles[i])

	for left < right {
		mid := left + (right - left) / 2
		if canFinishInTime(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left

}

func canFinishInTime(piles []int, h int, k int) bool {
	hrs := 0

	for _, p := range piles {
		hrs += (p + k - 1) / k
	}

	return hrs <= h
}
