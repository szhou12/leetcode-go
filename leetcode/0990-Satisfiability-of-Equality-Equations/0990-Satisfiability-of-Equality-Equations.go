package leetcode

func equationsPossible(equations []string) bool {
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < 26; i++ {
		uf.father[i] = i
	}

	for _, eqn := range equations {
		if eqn[1] == '=' {
			a, b := int(eqn[0]-'a'), int(eqn[3]-'a')
			if uf.Find(a) != uf.Find(b) {
				uf.Union(a, b)
			}
		}
	}

	for _, eqn := range equations {
		if eqn[1] == '!' {
			a, b := int(eqn[0]-'a'), int(eqn[3]-'a')
			if uf.Find(a) == uf.Find(b) {
				return false
			}
		}
	}

	return true
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
