package leetcode

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

/************* DFS ****************/
func numIslandsDFS(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	// corner case
	if m == 0 || n == 0 {
		return 0
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, &visited, i, j)
				res++
			}
		}
	}

	return res
}

func dfs(grid [][]byte, visited *[][]bool, x int, y int) {
	// base cases
	// case 1: out of board
	if !inBoard(grid, x, y) {
		return
	}
	// case 2: (x, y) is '0'/ocean
	if grid[x][y] == '0' {
		return
	}
	// case 3: (x, y) already visited
	if (*visited)[x][y] {
		return
	}

	// 当前层: mark (x,y) visited
	(*visited)[x][y] = true

	// DFS at four directions
	// Recursion
	for i := 0; i < 4; i++ {
		dx := x + dir[i][0]
		dy := y + dir[i][1]
		dfs(grid, visited, dx, dy)
	}
}

func inBoard(grid [][]byte, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}

/************* BFS ****************/
func numIslandsBFS(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	// corner case
	if m == 0 || n == 0 {
		return 0
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				bfs(grid, &visited, i, j)
				res++
			}
		}
	}

	return res
}

func bfs(grid [][]byte, visited *[][]bool, x int, y int) {
	queue := [][]int{{x, y}}
	(*visited)[x][y] = true

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				dx := cur[0] + dir[i][0]
				dy := cur[1] + dir[i][1]
				// base cases
				if !inBoard(grid, dx, dy) {
					continue
				}
				if grid[dx][dy] == '0' {
					continue
				}
				if (*visited)[dx][dy] {
					continue
				}
				queue = append(queue, []int{dx, dy})
				(*visited)[dx][dy] = true
			}
		}
	}
}

/************* Union Find ****************/
func numIslandsUF(grid [][]byte) int {
	uf := UnionFind{}
	uf.Init()
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			idx := i*m + j
			uf.father[idx] = idx
		}
	}

	// Union
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			cur := i*m + j
			for _, d := range dir {
				dx := i + d[0]
				dy := j + d[1]
				if !(0 <= dx && dx < m && 0 <= dy && dy < n) {
					continue
				}
				if grid[dx][dy] == '0' {
					continue
				}
				nei := dx*m + dy
				if uf.Find(cur) != uf.Find(nei) {
					uf.Union(cur, nei)
				}

			}
		}
	}

	// Count all connected components
	family := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			root := uf.Find(i*m + j)
			family[root] = true
		}
	}

	return len(family)
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
