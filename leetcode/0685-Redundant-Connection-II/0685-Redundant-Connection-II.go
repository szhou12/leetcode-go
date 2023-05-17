package leetcode

import "reflect"

func findRedundantDirectedConnection(edges [][]int) []int {

	// Step 1: find the node who has two parents (in-degree = 2)
	parent := make(map[int]int)
	var cand1, cand2 []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		// if 'to' already assigned with a parent, we find another parent
		if _, ok := parent[to]; ok {
			cand1 = []int{from, to}
			cand2 = []int{parent[to], to}
			break
		} else {
			parent[to] = from
		}
	}

	// Step 2: Union Find - find the redundant edge
	n := len(edges)
	uf := UnionFind{}
	uf.Init()
	for i := 1; i <= n; i++ {
		uf.father[i] = i
	}
	for _, edge := range edges {
		if reflect.DeepEqual(edge, cand1) {
			continue
		}
		from, to := edge[0], edge[1]
		if uf.Find(from) != uf.Find(to) {
			uf.Union(from, to)
		} else {
			if len(cand2) == 0 { // no node has 2 parents bc redundant edge connects root
				return edge
			} else {
				return cand2
			}
		}
	}
	return cand1
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
