package leetcode

import "math"

func maxRunTime(n int, batteries []int) int64 {
	left := 0
	right := math.MaxInt / n // divided by n in case of integer overflow

	for left < right {
		// check left = 0, right = 1 to determine how to write mid properly
		// mid := total time to run n computers
		mid := right - (right-left)/2
		if checkOk(mid, n, batteries) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return int64(left)

}

func checkOk(time int, n int, batteries []int) bool {
	totalTime := 0
	for _, battery := range batteries {
		// NOTE: NEED to take min here
		// Because each battery will serve at most 'time' minutes; exceeding part is useless
		totalTime += min(time, battery)
		if totalTime >= time*n {
			return true
		}
	}
	return false
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
