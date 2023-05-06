package leetcode

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	for _, edge := range allowedSwaps {
		a, b := edge[0], edge[1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
	}

	// Note: Update family ONLY after finishing Union Find bc unions will change during the process
	// Note: family needs to use count instead of bool bc element may have > 1
	family := make(map[int]map[int]int) // {index : {element: count} }
	for i, element := range source {
		if _, ok := family[uf.Find(i)]; !ok {
			family[uf.Find(i)] = make(map[int]int)
		}
		family[uf.Find(i)][element]++
	}

	res := 0
	for i, element := range target {
		if _, ok := family[uf.Find(i)][element]; !ok {
			res++
		} else {
			family[uf.Find(i)][element]--
			if family[uf.Find(i)][element] < 0 {
				res++
			}
		}
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
