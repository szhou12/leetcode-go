package amz

import "math"

/*
Given an n-length array costs[] and an integer 'weeks'.
You are asked to split costs[] into k subarrays whre each subarray takes the max of all elements.
Return the minimum sum of all subarrays.

Example:
costs = [1000, 500, 2000, 8000, 1500], weeks = 3

Output: 9500

Explain: 
Split costs[] into 3 subarrays: [1000], [500], max([2000, 8000, 1500]) = 8000
Sum = 1000 + 500 + 8000 = 9500
*/

/*
DP
Definition:
dp[i][k] := min sum of first i elements in costs[] split into k subarrays

Base Cases:
1. dp[i][0] = 0 // no 'weeks' to distribute
2. dp[i][1] = max(costs[:i]) // all elements in one subarray == max of all elements
3. dp[i][k] = math.MaxInt if i < k // not enough elements to distribute over k weeks

Recurrence:
dp[i][k] = min(dp[j][k-1] + max(costs[j:i])) for k-1 <= j < i

Note: 
1. k-1 <= j bc we need enough j elements to distribute over k-1 weeks
2. dp[j][k-1] := min sum of first j elements (costs[0:j-1]) split into k-1 subarrays

Time complexity = O(k * n^2)
*/
func minInputsPerWeek(costs []int, weeks int) int {
	n := len(costs)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, weeks+1)
		for k := 0; k <= weeks; k++ {
			dp[i][k] = math.MaxInt
		}
	}

	// Base Cases
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
		if i >= 1 {
			dp[i][1] = maxSlice(costs[:i])
		}
		
	}

	for k := 2; k <= weeks; k++ {
		for i := 1; i <= n; i++ {
			// not enough elements to distribute over k weeks
			if i < k {
				continue
			}

			for j := k-1; j < i; j++ {
				if dp[j][k-1] == math.MaxInt {
					continue
				}
				curSubarrayMax := maxSlice(costs[j:i])
				dp[i][k] = min(dp[i][k], dp[j][k-1] + curSubarrayMax)
			}
		}
	}

	return dp[n][weeks]
}

func maxSlice(nums []int) int {
	res := -1
	for _, num := range nums {
		res = max(res, num)
	}
	return res
}
