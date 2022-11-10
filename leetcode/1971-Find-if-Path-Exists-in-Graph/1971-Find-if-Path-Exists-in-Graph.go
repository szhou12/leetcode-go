package leetcode

// DFS - 会超时
func validPath_DFS(n int, edges [][]int, source int, destination int) bool {
	visited := make(map[int]bool)

	visited[source] = true
	return dfs(n, edges, source, destination, &visited)
}

func dfs(n int, edges [][]int, source int, destination int, visited *map[int]bool) bool {
	// base case
	if source == destination {
		return true
	}

	(*visited)[source] = true
	for i := 0; i < len(edges); i++ {
		start, end := edges[i][0], edges[i][1]
		if start == source && !(*visited)[end] {
			if dfs(n, edges, end, destination, visited) {
				return true
			}
		} else if end == source && !(*visited)[start] {
			if dfs(n, edges, start, destination, visited) {
				return true
			}
		}
	}
	return false
}
