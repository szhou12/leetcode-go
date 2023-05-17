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
	fx := uf.father[x]
	fy := uf.father[y]
	if fx < fy { // 谁小谁成为祖宗
		uf.father[fy] = fx
	} else {
		uf.father[fx] = fy
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
	n := len(edges)
	uf := UnionFind{}
	uf.Init()

	// 一开始, 每个node没有认过祖宗(root), 自己先成为自己的祖宗(root)
	for i := 1; i <= n; i++ {
		uf.father[i] = i
	}

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]

		// 如果a, b 已经有同一个祖宗, 目前这个edge就是导致环的additional edge
		if uf.Find(a) == uf.Find(b) {
			return edge
		} else { // 联姻
			uf.Union(a, b)
		}
	}

	return []int{}
}
