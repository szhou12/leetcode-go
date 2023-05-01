package leetcode

// Union Find
type UnionFind struct {
	father map[int]int
}

func (uf *UnionFind) Init() {
	uf.father = make(map[int]int) // 当前节点的父节点 {key: 当前节点, value: 父节点}
}

// 经过路径压缩的优化
func (uf *UnionFind) Find(x int) int {
	// 如果没到根节点, 递归地找到根节点, 并把当前节点挂到根节点下面
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

func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := UnionFind{}
	uf.Init()

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if _, ok := uf.father[a]; !ok {
			uf.father[a] = a
		}
		if _, ok := uf.father[b]; !ok {
			uf.father[b] = b
		}
		if uf.Find(a) != uf.Find(b) {
			uf.Union(a, b)
		}
	}
	return uf.Find(source) == uf.Find(destination)
}

// DFS - 会超时
func validPath_DFS(n int, edges [][]int, source int, destination int) bool {
	visited := make(map[int]bool)

	visited[source] = true
	return dfs(n, edges, source, destination, &visited)
}

func dfs(n int, edges [][]int, source int, destination int, visited *map[int]bool) bool {
	// base case
	if source == destination {
		return true
	}

	(*visited)[source] = true
	for i := 0; i < len(edges); i++ {
		start, end := edges[i][0], edges[i][1]
		if start == source && !(*visited)[end] {
			if dfs(n, edges, end, destination, visited) {
				return true
			}
		} else if end == source && !(*visited)[start] {
			if dfs(n, edges, start, destination, visited) {
				return true
			}
		}
	}
	return false
}
