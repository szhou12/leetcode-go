package leetcode

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < n; i++ {
		uf.father[i] = i
	}
	uf.father[firstPerson] = 0

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	secretSet := make(map[int]bool)
	secretSet[0] = true
	secretSet[firstPerson] = true

	// process meetings that occur at the same time
	for i := 0; i < len(meetings); i++ {
		sameTimePersons := make(map[int]bool)

		var j int
		for j = i; j < len(meetings) && meetings[i][2] == meetings[j][2]; j++ {
			p1, p2 := meetings[j][0], meetings[j][1]
			sameTimePersons[p1] = true
			sameTimePersons[p2] = true
			// Union
			if uf.Find(p1) != uf.Find(p2) {
				uf.Union(p1, p2)
			}
		}

		// contaminate persons to know the secret
		for p, _ := range sameTimePersons {
			if uf.Find(p) == 0 {
				secretSet[p] = true
			} else { // if not know secret, detach this person from the family
				uf.father[p] = p
			}
		}
		i = j - 1 // DO NOT FORGET!
	}

	res := make([]int, 0)
	for p, _ := range secretSet {
		res = append(res, p)
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
