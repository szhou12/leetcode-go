package leetcode

import (
	"math"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)

	// STEP 0: put startTime, endTime, profit all together [[s, e, p], ...]
	jobs := make([][]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}

	// STEP 1: sort by endTime
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	res := math.MinInt
	// STEP 2: DP
	// dp[i] = [endTime of job i, max profit] = max profit at job i's endTime (may or may not include job i's profit)
	// dp[i].profit = max(dp[i-1].profit, dp[j].profit + job[i].profit) where job j is the latest job that is compatible with job i (i.e. job[j].endTime <= job[i].endTime)
	// dp[i].endTime = job[i].endTime
	dp := make([][]int, 0)
	// base case
	dp = append(dp, []int{jobs[0][1], jobs[0][2]})
	res = max(res, dp[0][1])
	// recurrence
	for i := 1; i < n; i++ {
		// case 1: not inluding job i's profit, directly inherit from dp[i-1]'s profit
		curProfit := dp[i-1][1]
		// case 2: including job i's profit
		// Use binary search (upperBound) to find job j
		j := upperBound(dp, jobs[i][0]) - 1 // -1 to get job j s.t. dp[j].endTime <= job[i].startTime
		if j < 0 {                          // out of bound, no previous compatible job
			curProfit = max(curProfit, jobs[i][2])
		} else {
			curProfit = max(curProfit, dp[j][1]+jobs[i][2])
		}
		dp = append(dp, []int{jobs[i][1], curProfit})

		res = max(res, curProfit)
	}

	return res
}

// find the first index i s.t. nums[i][0] > target
// nums = [[endTime, maxProfit], ...]
// target = startTime
func upperBound(nums [][]int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid][0] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left][0] > target {
		return left
	} else {
		return left + 1
	}
}
