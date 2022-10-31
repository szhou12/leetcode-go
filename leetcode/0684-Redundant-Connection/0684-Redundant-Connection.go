package leetcode

type UnionFind struct {
	father map[int]int
}

func (uf *UnionFind) Init() {
	uf.father = make(map[int]int)
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.father[x] {
		uf.father[x] = uf.Find(uf.father[x])
	}
	return uf.father[x]
}

func (uf *UnionFind) Union(x int, y int) {
	x = uf.father[x]
	y = uf.father[y]
	if x < y { // 谁小谁成为祖宗
		uf.father[y] = x
	} else {
		uf.father[x] = y
	}
}

// 这个写法更明确联姻是发生在祖宗辈 而不是input
// func (uf *UnionFind) Union(x int, y int) {
//     fx := uf.father[x]
//     fy := uf.father[y]
//     if fx < fy {
//         uf.father[fy] = fx
//     } else {
//         uf.father[fx] = fy
//     }
// }

func findRedundantConnection(edges [][]int) []int {
	uf := UnionFind{}
	uf.Init()

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]

		// 如果当前node没有认过祖宗(root), 自己先成为自己的祖宗(root)
		if _, ok := uf.father[a]; !ok {
			uf.father[a] = a
		}
		if _, ok := uf.father[b]; !ok {
			uf.father[b] = b
		}

		// 如果a, b 已经有同一个祖宗, 目前这个edge就是导致环的additional edge
		if uf.Find(a) == uf.Find(b) {
			return edge
		} else { // 联姻
			uf.Union(a, b)
		}
	}

	return []int{}
}
