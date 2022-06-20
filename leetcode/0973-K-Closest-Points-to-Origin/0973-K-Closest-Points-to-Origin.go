package leetcode

import "sort"

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		res[i] = points[i]
	}

	return res
}
