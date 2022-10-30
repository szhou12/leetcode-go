package leetcode

import "sort"

// DP - O(6n*6n)
func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	cubes := make([][]int, 0)
	for i := 0; i < n; i++ {
		l, w, h := cuboids[i][0], cuboids[i][1], cuboids[i][2]
		cubes = append(cubes, []int{l, w, h, i})
		cubes = append(cubes, []int{l, h, w, i})
		cubes = append(cubes, []int{w, l, h, i})
		cubes = append(cubes, []int{w, h, l, i})
		cubes = append(cubes, []int{h, l, w, i})
		cubes = append(cubes, []int{h, w, l, i})
	}
	sort.Slice(cubes, func(i, j int) bool {
		if cubes[i][0] != cubes[j][0] {
			return cubes[i][0] < cubes[j][0]
		}
		if cubes[i][1] != cubes[j][1] {
			return cubes[i][1] < cubes[j][1]
		}
		if cubes[i][2] != cubes[j][2] {
			return cubes[i][2] < cubes[j][2]
		}
		return cubes[i][3] < cubes[j][3]
	})

	dp := make([]int, 6*n)
	// Base cases
	for i := 0; i < 6*n; i++ {
		dp[i] = cubes[i][2] // i-th cube height itself
	}
	// Recurrence
	for i := 1; i < 6*n; i++ {
		for j := 0; j < i; j++ { // 注意: 是 cube_j <= cube_i
			if cubes[j][0] <= cubes[i][0] && cubes[j][1] <= cubes[i][1] && cubes[j][2] <= cubes[i][2] && cubes[j][3] != cubes[i][3] {
				dp[i] = max(dp[i], dp[j]+cubes[i][2])
			}
		}
	}

	res := 0
	for i := 0; i < 6*n; i++ {
		res = max(dp[i], res)
	}

	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP - O(n*n)
// 每个cuboid只考虑一种变体: (min(l, w, h), mid(l, w, h), max(l, w, h))
func maxHeight_optimal(cuboids [][]int) int {
	n := len(cuboids)

	// 每个cuboid 内部升序排列, 得到变体(min(l, w, h), mid(l, w, h), max(l, w, h))
	for i := 0; i < n; i++ {
		sort.Ints(cuboids[i])
	}
	// cuboids之间再按照维度依次升序排列
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}
