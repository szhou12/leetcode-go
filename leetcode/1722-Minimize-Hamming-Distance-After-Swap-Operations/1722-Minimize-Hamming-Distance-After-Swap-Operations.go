package leetcode

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}
}

type UnionFind struct {
	father map[int]int
}

func (uf *UnionFind) Init() {
	uf.father = make(map[int]int)
}

func (uf *UnionFind) Union(X int, y int) {
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
