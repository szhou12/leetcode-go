package amz

import "sort"

/*
Leetcode 0056. Merge Intervals

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example:
intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]

intervals = [[1,4],[4,5]]
Output: [[1,5]]
*/

func mergeIntervals(intervals [][]int) [][]int {
	// Sort intervals by start time in ascending order
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	// dynamically update the last interval in res by updating its ending time to achieve merging if necessary
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] { // merge if overlapping: start-time_i <= end-time_(i-1)
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
