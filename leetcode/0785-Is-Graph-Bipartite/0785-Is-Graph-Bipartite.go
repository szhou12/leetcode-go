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
// Two colors: 0, 1
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

// Solution 2: Union Find - 相邻的节点不能被Union在同一个群组里
func isBipartite_UnionFind(graph [][]int) bool {
	n := len(graph)
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	for i, neis := range graph {
		// Note: neis may have length = 0
		var first int
		if len(neis) > 0 {
			first = neis[0] // prepare for union neis
		}

		for _, nei := range neis {
			// if i-th node and any of its neighbor are in the same group, it means NO bipartite
			if uf.Find(i) == uf.Find(nei) {
				return false
			}
			uf.Union(first, nei)
		}

	}
	return true

}

type UnionFind struct {
	father map[int]int
}

func (uf *UnionFind) Init() {
	uf.father = make(map[int]int)
}

func (uf *UnionFind) Union(x int, y int) {
	fx := uf.father[x]
	fy := uf.father[y]
	if fx < fy {
		uf.father[fy] = fx
	} else {
		uf.father[fx] = fy
	}
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.father[x] {
		uf.father[x] = uf.Find(uf.father[x])
	}
	return uf.father[x]
}
