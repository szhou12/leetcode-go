package leetcode

import "sort"

// for each size (ascending order) of boxes[j], find the range it can cover in packages,
// if packages can't be fully covered, then -1
// if packages can be fully covered, find the min wasted sum
// use prefix-sum to calculate the sum of wasted space in covered range.
// try out all j's to find the minimun
func minWastedSpace(packages []int, boxes [][]int) int {
	n := len(packages)
	sort.Ints(packages) // O(nlogn)

	// prefix sum of packages
	prefixSum := make([]int, n)
	prefixSum[0] = packages[0]
	for k := 1; k < n; k++ { // O(n)
		prefixSum[k] = prefixSum[k-1] + packages[k]
	}

	res := -1

	// total TC = O(k * (mlogm + mlogn)) = O(km * logn)
	for _, boxSlice := range boxes { // O(k)
		sort.Ints(boxSlice) // O(mlogm)

		// edge case: if the largest box cannot fit the largest package, skip this box slice
		if boxSlice[len(boxSlice)-1] < packages[n-1] {
			continue
		}

		// for each box size from smallest to largest, find the upper bound of packages it can cover
		// left, right are packages indices
		// cover range = [left+1, right]
		left := -1
		curSum := 0
		for i := 0; i < len(boxSlice); i++ { // O(m)
			right := upperBound(packages, boxSlice[i]) - 1 // O(logn)

			// edge case: no upper bound found
			if right < 0 {
				continue
			}

			if left != -1 {
				curSum += (right-left)*boxSlice[i] - (prefixSum[right] - prefixSum[left])
			} else {
				curSum += (right-left)*boxSlice[i] - (prefixSum[right])
			}

			// update left
			left = right
			// early stop: if last package has already been covered, break the loop
			if left == n-1 {
				break
			}
		}

		if res == -1 {
			res = curSum
		} else {
			res = min(res, curSum)
		}

	}

	return res % int(1e9+7)
}

func upperBound(nums []int, target int) int {
	res := sort.Search(
		len(nums),
		func(i int) bool {
			return nums[i] > target
		},
	)

	return res
}
