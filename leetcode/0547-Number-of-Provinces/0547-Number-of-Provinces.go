package leetcode

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && isConnected[i][j] == 1 {
				if uf.Find(i) != uf.Find(j) {
					uf.Union(i, j)
				}
			}
		}
	}

	family := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		family[root] = true
	}

	return len(family)
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
