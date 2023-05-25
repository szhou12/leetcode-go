package leetcode

import "sort"

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	// build a well at node i == connect node i to node 0 (imaginary source node) with cost wells[i]
	for i, cost := range wells {
		pipes = append(pipes, []int{0, i + 1, cost})
	}

	uf := UnionFind{}
	uf.Init()
	for i := 0; i <= n; i++ {
		uf.father[i] = i
	}

	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})

	totalCost := 0
	for _, edge := range pipes {
		a, b, cost := edge[0], edge[1], edge[2]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
			totalCost += cost
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
