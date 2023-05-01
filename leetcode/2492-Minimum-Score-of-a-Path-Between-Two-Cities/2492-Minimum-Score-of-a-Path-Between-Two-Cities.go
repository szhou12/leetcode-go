package leetcode

import "math"

func minScore(n int, roads [][]int) int {
	uf := UnionFind{}
	uf.Init()

	// Step 1: Init all nodes ancestor to themselves
	for i := 1; i <= n; i++ {
		uf.father[i] = i
	}

	// Step 2: Union connected nodes
	for _, road := range roads {
		a, b := road[0], road[1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
	}

	// Step 3: Check each road, Find min distance between node i and node 1
	// NOTE: if node i is connected with node 1, node i must be connected with node n
	// i.e. node i belongs to the connected graph (set) that contains node 1 & n
	// Bc input condition: There is at least one path between 1 and n
	res := math.MaxInt
	for _, road := range roads {
		a, dist := road[0], road[2]
		if uf.Find(a) == uf.Find(1) {
			res = min(res, dist)
		}
	}

	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
