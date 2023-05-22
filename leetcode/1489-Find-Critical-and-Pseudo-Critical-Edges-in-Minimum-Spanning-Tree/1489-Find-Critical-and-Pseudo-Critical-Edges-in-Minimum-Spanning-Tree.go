package leetcode

import (
	"math"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// Step 0: add origianl index to each edge
	for i := 0; i < len(edges); i++ {
		edges[i] = append(edges[i], i)
	}

	// Step 1: sort edges by weight + get global MST weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	minMST := kruskal(n, edges, -1)

	// Step 2: find critical edge index
	critical := make(map[int]bool) // { edge index : true }
	for i := 0; i < len(edges); i++ {
		mst := kruskal(n, edges, i)
		if mst > minMST {
			critical[edges[i][3]] = true
		}
	}

	// Step 3: find pseudo-critical edge index
	pseudo := make(map[int]bool) // { edge index : true }
	for i := 0; i < len(edges); i++ {
		if critical[edges[i][3]] { // skip critical edge bc critical edges and pseudo-critical edges are mutually exclusive
			continue
		}
		// prepend cur edge to the front to force it to be included in MST construction
		edge := edges[i]
		edges = append([][]int{edge}, edges...)
		mst := kruskal(n, edges, -1)
		if mst == minMST {
			pseudo[edge[3]] = true
		}
		// recover edges
		edges = edges[1:]
	}

	// Step 4: convert map to slice
	res1 := make([]int, 0)
	for k := range critical {
		res1 = append(res1, k)
	}
	res2 := make([]int, 0)
	for k := range pseudo {
		res2 = append(res2, k)
	}

	return [][]int{res1, res2}
}

func kruskal(n int, edges [][]int, skipIdx int) int {
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	totalWeight := 0
	for i, edge := range edges {
		if i == skipIdx {
			continue
		}
		a, b, weight := edge[0], edge[1], edge[2]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
			totalWeight += weight
		}
	}

	// check if all nodes are connected
	for i := 0; i < n; i++ {
		if uf.Find(i) != uf.Find(0) {
			return math.MaxInt
		}
	}
	return totalWeight
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
