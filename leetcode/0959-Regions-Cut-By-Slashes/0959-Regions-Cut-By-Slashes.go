package leetcode

func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := UnionFind{}
	uf.Init()

	// Step 1: 外围点聚为一类，内部点各自为政
	// NOTE: n-len has [0, ..., n] points, so 1-D index = i*(n+1) + j
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			idx := i*(n+1) + j

			uf.father[idx] = idx

			if i == 0 || j == 0 || i == n || j == n {
				uf.father[idx] = 0
			}
		}
	}

	// Step 2: iterate slash
	count := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == ' ' {
				continue
			}

			var a, b int
			if grid[i][j] == '/' {
				a = i*(n+1) + (j + 1)
				b = (i+1)*(n+1) + j
			} else if grid[i][j] == '\\' {
				a = i*(n+1) + j
				b = (i+1)*(n+1) + (j + 1)
			}

			if uf.Find(a) == uf.Find(b) {
				count++
			} else {
				uf.Union(a, b)
			}
		}
	}

	return count

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
