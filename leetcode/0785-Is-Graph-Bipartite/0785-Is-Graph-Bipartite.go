package leetcode

func isBipartite(graph [][]int) bool {
	visited := make(map[int]int)

	// loop through every node as the starting node just in case there exists an isolated node (disconnected)
	for node, _ := range graph {
		if !bfs(node, visited, graph) {
			return false
		}
	}
	return true
}

// BFS returns true if all connected nodes to the starting node can be safely colored bipartitely
// BFS returns false if there is any conflicts
func bfs(node int, visited map[int]int, graph [][]int) bool {
	// check if the input node has been visited
	// if already visited, safe
	if _, ok := visited[node]; ok {
		return true
	}

	visited[node] = 0 // coloring input node as 0

	queue := make([]int, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		curNodeColor := visited[curNode]
		neiNodeColor := 0
		if curNodeColor == 0 {
			neiNodeColor = 1
		}

		for _, nei := range graph[curNode] {
			if _, ok := visited[nei]; !ok {
				visited[nei] = neiNodeColor
				queue = append(queue, nei)
			} else if visited[nei] != neiNodeColor {
				return false
			}
		}
	}
	return true
}
