package leetcode

import "math"

var TOTALCOLORS = 61

func isPrintable(targetGrid [][]int) bool {
	m, n := len(targetGrid), len(targetGrid[0])
	// Step 1: Construct directed graph on colors
	// 1a. find each color's upper-, lower-, left-, right-bound
	top := make([]int, TOTALCOLORS)
	bottom := make([]int, TOTALCOLORS)
	left := make([]int, TOTALCOLORS)
	right := make([]int, TOTALCOLORS)
	for i := 0; i < TOTALCOLORS; i++ {
		top[i] = math.MaxInt
		left[i] = math.MaxInt
		bottom[i] = math.MinInt
		right[i] = math.MinInt
	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			color := targetGrid[x][y]
			top[color] = min(top[color], x)
			bottom[color] = max(top[color], x)
			left[color] = min(left[color], y)
			right[color] = max(right[color], y)
		}
	}
	// 1b. construct adj-list repr; a node = one color
	next := make([][]int, TOTALCOLORS)
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			to := targetGrid[x][y]
			for c := 1; c < TOTALCOLORS; c++ {
				if top[c] <= x && x <= bottom[c] && left[c] <= y && y <= right[c] {
					if c != to {
						next[c] = append(next[c], to)
					}
				}
			}
		}
	}
	// 1c. calc degree
	degree := make([]int, TOTALCOLORS)
	for c := 1; c < TOTALCOLORS; c++ {
		for _, to := range next[c] {
			degree[to]++
		}
	}

	// Step 2: Topological Sort
	queue := make([]int, 0)
	// Start nodes: degree == 0
	for c := 1; c < TOTALCOLORS; c++ {
		if degree[c] == 0 {
			queue = append(queue, c)
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for _, nei := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Check if has a cycle
	for c := 1; c < TOTALCOLORS; c++ {
		if degree[c] != 0 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
