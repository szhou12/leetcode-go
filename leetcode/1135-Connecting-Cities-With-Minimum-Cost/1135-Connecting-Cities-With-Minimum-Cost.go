package leetcode

import "sort"

func minimumCost(n int, connections [][]int) int {
	uf := UnionFind{}
	uf.Init()
	for i := 1; i <= n; i++ {
		uf.father[i] = i
	}

	// Sort edges in increasing order
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	totalCost := 0
	for _, edge := range connections {
		a, b, cost := edge[0], edge[1], edge[2]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
			totalCost += cost
		}
	}

	// check if all nodes connected
	for i := 1; i <= n; i++ {
		if uf.Find(i) != uf.Find(1) {
			return -1
		}
	}
	return totalCost
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
