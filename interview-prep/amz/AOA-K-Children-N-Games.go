package amz

import (
	"sort"
)

/*
Given k bins with identical capacities, n items. i-th item has volumes[i] in volumes array.
Each bin must contain at least one item but no more than two items. No item can be left unassigned.
Find the minimum capacity for each bin.

Example:
n = 4
volumes = [9, 2, 4, 6]
k = 3
Output: Min Capacity = 9
Explanation: Bin 1 = [9], Bin 2 = [2, 4], Bin 3 = [6]
*/

func findMinCapacity(k int, volumes []int) int {

	// sort volumes in descending order
	sort.Slice(volumes, func (i, j int) bool {
		return volumes[i] > volumes[j]
	})

	// Binary Search to guess a capacity
	left, right := volumes[0], volumes[0] + volumes[1]
	for left < right {
		mid := left + (right-left)/ 2
		if canFit(k, volumes, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left

}

func canFit(k int, volumes []int, capacity int) bool {
	n := len(volumes)
	left := 0
	right := n-1
	binsUsed := 0

	// if every bin containing 2 items is uable to cover all n items, we are not able to dispense all items
	if 2*k < n {
		return false
	}

	bins := make([][]int, k)

	// Assign largest k items to each of k bins
	for i := 0; i < k && left <= right; i++ {
		bins[i] = []int{volumes[left]}
		left++
		binsUsed++
	}

	// Assign remaining items to bins one by one starting from the smallest item
	binIndex := 0
	for left <= right {
		if binIndex >= k {
			return false
		}

		if len(bins[binIndex]) == 2 {
			binIndex++
			continue
		}

		if bins[binIndex][0] + volumes[right] <= capacity {
			bins[binIndex] = append(bins[binIndex], volumes[right])
			right--
		} else {
			binIndex++
		}
	}

	if left <= right {
		return false
	}

	return binsUsed == k
}
