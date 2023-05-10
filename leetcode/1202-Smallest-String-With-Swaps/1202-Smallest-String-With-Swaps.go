package leetcode

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	// Step 1: Union Find to form connected componets
	for _, edge := range pairs {
		a, b := edge[0], edge[1]
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
	}

	family := make(map[int][]int) // { root index : member indices} NOTE: member indices MUST be increasing order
	for i := 0; i < n; i++ {
		root := uf.Find(i)
		if _, ok := family[root]; !ok {
			family[root] = make([]int, 0)
		}

		family[root] = append(family[root], i)
	}

	// Step 2: Replace sorted chars per component
	// string is immutable, so convert to slice of runes to replace char by index
	source := []rune(s)
	for _, component := range family {
		chars := make([]rune, 0)
		for _, idx := range component {
			chars = append(chars, rune(s[idx]))
		}
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		for _, idx := range component {
			source[idx] = chars[0]
			chars = chars[1:]
		}
	}

	return string(source)
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
