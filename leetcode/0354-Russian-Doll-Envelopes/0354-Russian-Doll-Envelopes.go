package leetcode

import "sort"

// Optimal Solution - Greedy + Binary Search
func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		// 注意！！！如果w相同的两个元素, 把h较大的排在前面
		// 这样在找LIS时，就不需要额外考虑当w相同时，后一个元素不能加进来
		// 这样排序后，就可以完全抛开w这个维度，专注在h这个维度上找LIS了
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// 维护一个数组，里面的元素单调递增
	// Case 1: 如果当前元素大于所有目前数组内的元素, 直接append
	// Case 2: 如果当前元素小于某些目前数组内的元素, 找到数组内第一个比它大的元素并用它替换掉
	res := make([]int, 0)
	for _, env := range envelopes { // O(n)
		if len(res) == 0 || res[len(res)-1] < env[1] {
			res = append(res, env[1])
		} else {
			i := findFirstLarger(res, env[1]) // O(logn)
			res[i] = env[1]
		}
	}

	return len(res)
}

func findFirstLarger(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// My Solution - DP - 会超时!!!
// step 1: sort by increasing order
// step 2:
// dp[i] = max # of envelopes we can Russian Doll ending at envelopes[i]
// Recurrence:
// dp[i] = max(dp[j] + 1) for 1 <= j < i and i's (w, h) > j's (w, h)
func maxEnvelopes_mysoln(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	n := len(envelopes)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
