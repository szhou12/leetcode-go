package leetcode

import "math"

func minimumArrayLength(nums []int) int {
	// Step 1: find min value and its frequency
	x, cnt := math.MaxInt, 0
	for _, num := range nums {
		if num < x {
			x, cnt = num, 1
		} else if num == x {
			cnt++
		}
	}

	// Strategy 1: if x occurs only once, all other values can be removed, return the length 1
	if cnt == 1 {
		return 1
	}

	// Strategy 2: if x occurs more than once, BUT there exists y s.t. y % x != 0, we can get a new min value bc y % x < x.
	// we return to the Strategy 1
	for _, y := range nums {
		if y % x != 0 {
			return 1
		}
	}

	// Strategy 3: if x occurs more than once, and all other values are divisible by x, we need to look at even/odd
	// Case 1: cnt = even (e.g. 2 2 2 2 => 0 0)
	// Case 2: cnt = odd (e.g. 4 4 4 => 0 4)
	if cnt %2 == 0 {
		return cnt/2
	} else {
		return cnt/2 + 1
	}
}
