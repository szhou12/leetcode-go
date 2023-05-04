package leetcode

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}

	res := make([]bool, 0)
	for _, req := range requests {
		p1, p2 := req[0], req[1]
		fatherP1, fatherP2 := uf.Find(p1), uf.Find(p2)

		becomeFriend := true
		for _, restrict := range restrictions {
			r1, r2 := restrict[0], restrict[1]
			fatherR1, fatherR2 := uf.Find(r1), uf.Find(r2)

			if (fatherR1 == fatherP1 && fatherR2 == fatherP2) || (fatherR1 == fatherP2 && fatherR2 == fatherP1) {
				becomeFriend = false
				break
			}
		}

		res = append(res, becomeFriend)
		if becomeFriend {
			uf.Union(p1, p2)
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
