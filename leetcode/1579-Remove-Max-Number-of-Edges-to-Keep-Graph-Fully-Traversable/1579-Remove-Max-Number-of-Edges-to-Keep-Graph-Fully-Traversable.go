package leetcode

func maxNumEdgesToRemove(n int, edges [][]int) int {
	// Step 0: separate 3 types of edges
	edgesBoth := make([][]int, 0)
	edgesAlice := make([][]int, 0)
	edgesBob := make([][]int, 0)
	for _, edge := range edges {
		edgeType, a, b := edge[0], edge[1], edge[2]
		if edgeType == 3 {
			edgesBoth = append(edgesBoth, []int{a, b})
		}
		if edgeType == 1 {
			edgesAlice = append(edgesAlice, []int{a, b})
		}
		if edgeType == 2 {
			edgesBob = append(edgesBob, []int{a, b})
		}
	}

	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	// Step 1: Prioritize type-3 edges that both Alice and Bob can use
	countBoth := 0
	for _, edge := range edgesBoth {
		a, b := edge[0], edge[1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
			countBoth++
		}
	}

	// Step 2: 补上 type-1 edges for Alice
	countAlice := connectGraph(uf, edgesAlice)
	// if Alice can't connect all nodes, return -1
	if countBoth+countAlice < n-1 {
		return -1
	}

	// Step 3: 补上 type-2 edges for Bob
	countBob := connectGraph(uf, edgesBob)
	// if Bob can't connect all nodes, return -1
	if countBoth+countBob < n-1 {
		return -1
	}

	// return max # of edges that can be removed
	return len(edges) - (countBoth + countAlice + countBob)

}

// Input:
//  - graph already connected by type-3 edges
//  - edges that ONLY current person can use
func connectGraph(common UnionFind, edges [][]int) int {
	// make a copy of common to avoid direct change
	uf := UnionFind{}
	uf.Init()
	for k, v := range common.father {
		uf.father[k] = v
	}

	count := 0
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
			count++
		}
	}

	return count
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
