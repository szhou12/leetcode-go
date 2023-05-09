package leetcode

func makeConnected(n int, connections [][]int) int {
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	// Step 1: Union connected components
	degree := make([]int, n) // in-degree of every node
	for _, edge := range connections {
		a, b := edge[0], edge[1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
		degree[a]++
		degree[b]++
	}

	// Step 2: count # of edges & # of nodes for each component
	// NOTE: each edge is double counted from both ends
	edges := make(map[int]int)
	nodes := make(map[int]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		edges[root] += degree[i]
		nodes[root]++
	}

	// Step 3: count # of redundant edges for each component
	redundant := 0
	for root, cables := range edges {
		redundant += cables/2 - (nodes[root] - 1)
	}

	components := len(edges)
	if redundant >= components-1 {
		return components - 1
	}
	return -1
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
