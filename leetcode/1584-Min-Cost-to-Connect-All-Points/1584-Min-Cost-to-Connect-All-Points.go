package leetcode

import "sort"

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	// Step 1: Calculate edge weight O(n^2)
	edges := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xi, yi := points[i][0], points[i][1]
			xj, yj := points[j][0], points[j][1]
			weight := abs(xi-xj) + abs(yi-yj)
			edges = append(edges, []int{weight, i, j})
		}
	}

	// Step 2: Kruskal Algorithm
	res := kruskal(n, edges)
	return res
}

func kruskal(n int, edges [][]int) int {
	// UnionFind stores each point's index
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	// Step 1: Sort edges in increasing order
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	// Step 2: Union Find
	totalWeight := 0
	for _, edge := range edges {
		weight, a, b := edge[0], edge[1], edge[2]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
			totalWeight += weight
		}
	}

	return totalWeight
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
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
