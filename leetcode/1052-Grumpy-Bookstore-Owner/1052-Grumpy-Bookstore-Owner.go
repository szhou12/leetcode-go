package leetcode

// Sliding Window Standard Solution
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	sum := 0
	// 先把基本盘加进去: grumpy == 0 的地方都是基本盘, 因为不管在不在window里, 都要考虑进result
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}

	res := 0
	// sliding window
	// 注意：这里 sliding window 从 index=0 滑入
	// 吃是从index=0开始，吐是从可以吐的时候才开始吐
	// sliding window 不用非得从 [0, 0+minutes]处才开始
	// 因为一开始没有吃满window size个的情况, 每多吃一个, sum是non-decreasing的, 所以能保证正确性
	for right := 0; right < len(customers); right++ {
		// 吃
		if grumpy[right] == 1 {
			sum += customers[right]
		}
		// 吐
		if right-minutes >= 0 && grumpy[right] == 1 {
			sum -= customers[right]
		}
		res = max(res, sum)
	}
	return res
}

// My Implementation of Sliding Window
func maxSatisfied_mysoln(customers []int, grumpy []int, minutes int) int {
	// initial #: 0 - 0+minutes
	sum := 0
	for i := 0; i < len(customers); i++ {
		if 0 <= i && i < minutes {
			sum += customers[i]
		} else if grumpy[i] == 0 {
			sum += customers[i]
		}
	}
	res := sum
	for l := 1; l+minutes <= len(customers); l++ {
		r := l + minutes - 1
		if grumpy[r] == 1 {
			sum += customers[r]
		}
		if grumpy[l-1] == 1 {
			sum -= customers[l-1]
		}
		res = max(res, sum)
	}
	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
