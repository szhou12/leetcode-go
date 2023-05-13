package leetcode

func removeStones(stones [][]int) int {
	uf := UnionFind{}
	uf.Init()
	n := 10000
	mapX := make(map[int][]int) // { x-coord : [1-D index]}
	mapY := make(map[int][]int) // { y-coord : [1-D index]}

	// Step 1: same X/Y-coord points grouped together
	for _, s := range stones {
		i, j := s[0], s[1]
		idx := i*n + j
		uf.father[idx] = idx

		if _, ok := mapX[i]; !ok {
			mapX[i] = make([]int, 0)
		}
		mapX[i] = append(mapX[i], idx)

		if _, ok := mapY[j]; !ok {
			mapY[j] = make([]int, 0)
		}
		mapY[j] = append(mapY[j], idx)
	}

	// Union by x-coord
	for _, ids := range mapX {
		firstId := ids[0]
		for i := 1; i < len(ids); i++ {
			if uf.Find(firstId) != uf.Find(ids[i]) {
				uf.Union(firstId, ids[i])
			}
		}
	}
	// Union by y-coord
	for _, ids := range mapY {
		firstId := ids[0]
		for i := 1; i < len(ids); i++ {
			if uf.Find(firstId) != uf.Find(ids[i]) {
				uf.Union(firstId, ids[i])
			}
		}
	}

	// Step 2: count # of unified groups
	family := make(map[int]bool) // {ancestor : true}
	for _, s := range stones {
		i, j := s[0], s[1]
		root := uf.Find(i*n + j)
		family[root] = true
	}

	return len(stones) - len(family)

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
