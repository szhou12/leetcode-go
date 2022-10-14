package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// Step 1: Sort by END time in increasing order!!! if we want to loop from left to right
	sort.Slice(intervals, func(i, j int) bool {
		// if block 其实没有必要，删掉也能保证正确性,
		// 因为不管是保留start较早的还是较晚的，只要他们不与之后的时间段有重叠, 删去任何一个都是正确的
		// if intervals[i][1] == intervals[j][1] {
		// 	return intervals[i][0] < intervals[j][0]
		// }
		return intervals[i][1] < intervals[j][1]
	})

	// Step 2: Loop from left to right; check whether 'end' falls within [start_i, end_i]
	cnt := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end && end <= intervals[i][1] { // overlapped
			cnt++
		} else {
			end = intervals[i][1]
		}
	}
	return cnt
}
