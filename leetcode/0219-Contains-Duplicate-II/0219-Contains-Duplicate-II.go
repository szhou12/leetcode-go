package leetcode

import "math"

type item struct {
	lastIndex int
	distance  int
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]item)

	for idx, num := range nums {
		if it, ok := dict[num]; ok {
			newDist := min(it.distance, abs(it.lastIndex, idx))
			if newDist <= k {
				return true
			}
			newItem := item{lastIndex: idx, distance: newDist}
			dict[num] = newItem
		} else {
			newItem := item{lastIndex: idx, distance: math.MaxInt}
			dict[num] = newItem
		}
	}

	return false
}

func abs(a int, b int) int {
	dist := a - b
	if dist < 0 {
		return -dist
	}
	return dist
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
