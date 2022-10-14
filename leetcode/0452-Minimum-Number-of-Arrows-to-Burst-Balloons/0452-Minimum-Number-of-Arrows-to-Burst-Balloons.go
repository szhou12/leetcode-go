package leetcode

import "sort"

// Greedy Algorithm - sort by end in increasing order
func findMinArrowShots(points [][]int) int {
	// Step 1: sort by end in increasing order
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// Step 2: greedy - if no overlap, then we need one more arrow
	cnts := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end && end <= points[i][1] {
			continue
		} else {
			cnts++
			end = points[i][1]
		}
	}
	return cnts
}

//  i
// (1,6) (2,8) (7,12), (10,16)
//    e
