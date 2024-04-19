package leetcode

import "math"

func repairCars(ranks []int, cars int) int64 {
	left, right := 0, math.MaxInt
	for left < right {
		mid := left + (right-left)/2
		if canRepair(mid, ranks, cars) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}

func canRepair(time int, ranks []int, cars int) bool {
	maxRepairs := 0
	for _, r := range ranks {
		maxRepairs += int(math.Sqrt(float64(time / r)))
	}
	if maxRepairs >= cars {
		return true
	}
	return false
}
