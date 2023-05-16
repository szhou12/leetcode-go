package leetcode

func minSwapsCouples(row []int) int {
	n := len(row)
	uf := UnionFind{}
	uf.Init()

	// Union couples
	for i := 0; i < n; i += 2 {
		uf.father[i] = i
		uf.father[i+1] = i
	}

	// Union seats
	for i := 0; i < n; i += 2 {
		a, b := row[i], row[i+1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
	}

	// Count # of people in each component
	size := make(map[int]int)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		size[root]++
	}

	// Count # of swaps within each component = # of couples - 1
	res := 0
	for _, people := range size {
		couples := people / 2
		res += couples - 1
	}

	return res

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
