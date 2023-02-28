package leetcode

import "math"

// left[i] := length of longest LIS in nums[0:i] ending at i
// right[i] := length of longest LIS in nums[i:n-1] ending at i
func minimumMountainRemovals(nums []int) int {
	n := len(nums)

	left := make([]int, n)
	// base cases
	for i := 0; i < n; i++ {
		left[i] = 1
	}
	// recurrence
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				left[i] = max(left[i], left[j]+1)
			}
		}
	}

	right := make([]int, n)
	// base cases
	for i := n - 1; i >= 0; i-- {
		right[i] = 1
	}
	// recurrence
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				right[i] = max(right[i], right[j]+1)
			}
		}
	}

	res := n
	for i := 0; i < n; i++ {
		if left[i] >= 2 && right[i] >= 2 { // i左手边必须至少有一个元素, 加上i自己就是必须有至少两个元素; 右手边同理
			res = min(res, n-(left[i]+right[i]-1))
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 这个解法left[], right[]的物理意义比较直接, 但是实现起来在算保留元素个数的时候有点麻烦, 不如按LIS的物理意义定义来的直观
// left[i] := min # removals to make nums[0:i] ending at i strictly increasing
// right[i] := min # removals to make nums[i:n-1] ending at i strictly decreasing
// base cases:
// left[0] = 0
// left[1..n] = i 只保留ith元素，前面的全删了, 这样不可以
// left[i] = min(left[j] + ((i-1)-(j+1)+1)) for 0 <= j < i and nums[i] > nums[j]
//   100, 92, 89, 77, 74, 66, 64, 66, 64
// l  0    1  2   3   4   5    6   6   8
// r  2    2  2   2   2   2    2   0   0
func minimumMountainRemovals_soln2(nums []int) int {
	n := len(nums)

	left := make([]int, n)
	// base cases
	left[0] = 0
	for i := 1; i < n; i++ {
		left[i] = i
	}
	// recurrence
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				left[i] = min(left[i], left[j]+((i-1)-(j+1)+1))
			}
		}
	}

	right := make([]int, n)
	// base cases
	right[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		right[i] = (n - 1) - (i + 1) + 1
	}
	// recurrence
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				right[i] = min(right[i], right[j]+((j-1)-(i+1)+1))
			}
		}
	}

	res := math.MaxInt
	for i := 1; i < n-1; i++ {
		// 这里的if必须有, 不能发生把 ith元素 之前/之后的元素全删了情况, 至少要保留一个, 否则不构成mountain array
		if ((i-0+1)-left[i] >= 2) && (((n-1)-i+1)-right[i] >= 2) {
			res = min(res, left[i]+right[i])
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
