package leetcode

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < 26; i++ {
		uf.father[i] = i
	}

	n := len(s1)
	for i := 0; i < n; i++ {
		a, b := int(s1[i]-'a'), int(s2[i]-'a')
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
	}

	res := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		root := uf.Find(int(baseStr[i] - 'a'))
		res[i] = byte(root + 'a')
	}
	return string(res)
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
