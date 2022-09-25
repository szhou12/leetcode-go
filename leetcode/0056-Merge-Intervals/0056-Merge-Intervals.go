package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	// MUST sort by starting time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] { // end time_(i-1) >= start time_(i)
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
