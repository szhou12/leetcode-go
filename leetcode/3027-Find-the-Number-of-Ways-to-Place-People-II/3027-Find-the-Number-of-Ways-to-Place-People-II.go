package leetcode

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	n := len(points)

	// Sort: sort by x increasing order, then by y decreasing order (do you know why?)
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	pairs := 0
	for i := 0; i < n; i++ { // Alice
		upper := points[i][1] // 上限
		lower := math.MinInt  // 下限: Alice和Bob之间的点中最大的y坐标
		for j := i + 1; j < n; j++ { // Bob
			if points[j][1] > upper {
				continue
			}
			if lower < points[j][1] && points[j][1] <= upper {
				pairs++
			}
		}
	}
	return pairs
}
