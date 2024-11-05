package leetcode

import "math"

func minTimeToReach(moveTime [][]int) int {
	dirs := [][]int{
		{0, -1},
		{-1, 0},
		{1, 0},
		{0, 1},
	}

	n, m := len(moveTime), len(moveTime[0])
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, m)
		for j := 0; j < m; j++ {
			visited[i][j] = math.MaxInt
		}
	}

	// dfs(x, y, curTime) := taking curTime to go from (0, 0) to (x, y)
	var dfs func(x int, y int, curTime int)
	dfs = func(x int, y int, curTime int) {
		// base case: If reached the destination
		if x == n-1 && y == m-1 {
			visited[x][y] = min(visited[x][y], curTime)
			return
		}

		// if we've been here earlier, no need to proceed
		if visited[x][y] <= curTime {
			return
		}

		// udpate visiting time at (x, y)
		visited[x][y] = curTime

		for _, dir := range dirs {
			dx, dy := x+dir[0], y+dir[1]
			if !(0 <= dx && dx < n && 0 <= dy && dy < m) {
				continue
			}

			nextTime := 0
			if curTime < moveTime[dx][dy] {
				nextTime = moveTime[dx][dy] + 1
			} else {
				nextTime = curTime + 1
			}

			dfs(dx, dy, nextTime)
		}
	}

	dfs(0, 0, 0)
	return visited[n-1][m-1]
}
