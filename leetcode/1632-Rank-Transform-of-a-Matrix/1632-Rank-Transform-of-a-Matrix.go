package leetcode

import "sort"

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	// UnionFind stores 1-D index
	uf := UnionFind{}
	uf.Init()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			uf.father[i*n+j] = i*n + j
		}
	}

	next := make([][]int, m*n)
	for i := 0; i < len(next); i++ {
		next[i] = make([]int, 0)
	}
	degree := make([]int, m*n)
	// Step 1: Reconstruct adj-list repr + Calculate degree + Union same-value elements
	// Per Row
	for i := 0; i < m; i++ {
		arr := make([][]int, 0) // [ [value, 1-D index], ... ]
		for j := 0; j < n; j++ {
			arr = append(arr, []int{matrix[i][j], i*n + j})
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i][0] < arr[j][0]
		})

		for k := 1; k < len(arr); k++ {
			// if same value, union(k-1, k)
			// if not, k-1 -> k
			if arr[k-1][0] == arr[k][0] {
				if uf.Find(arr[k-1][1]) != uf.Find(arr[k][1]) {
					uf.Union(arr[k-1][1], arr[k][1])
				}
			} else {
				next[arr[k-1][1]] = append(next[arr[k-1][1]], arr[k][1])
				degree[arr[k][1]]++
			}
		}
	}
	// Per Column
	for j := 0; j < n; j++ {
		arr := make([][]int, 0) // [ [value, 1-D index], ... ]
		for i := 0; i < m; i++ {
			arr = append(arr, []int{matrix[i][j], i*n + j})
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i][0] < arr[j][0]
		})

		for k := 1; k < len(arr); k++ {
			// if same value, union(k-1, k)
			// if not, k-1 -> k
			if arr[k-1][0] == arr[k][0] {
				if uf.Find(arr[k-1][1]) != uf.Find(arr[k][1]) {
					uf.Union(arr[k-1][1], arr[k][1])
				}
			} else {
				next[arr[k-1][1]] = append(next[arr[k-1][1]], arr[k][1])
				degree[arr[k][1]]++
			}
		}
	}

	// Step 2: Find all members for each group + all members degree added to family ancestor
	family := make(map[int][]int) // { family ancestor : [all members] }
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			root := uf.Find(i*n + j)
			if _, ok := family[root]; !ok {
				family[root] = make([]int, 0)
			}
			family[root] = append(family[root], i*n+j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			root := uf.Find(i*n + j)
			if root != i*n+j {
				degree[root] += degree[i*n+j]
			}
		}
	}

	// Step 3: Topological Sort - ONLY family ancestors enter queue
	queue := make([]int, 0)
	rank := 1
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	// Start nodes: family ancestor & degree == 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if uf.Find(i*n+j) == i*n+j && degree[i*n+j] == 0 {
				queue = append(queue, i*n+j)
			}
		}
	}
	// Loop
	for len(queue) > 0 {
		// Need to use size to segregate each round of loop bc we need to increment rank after each round
		size := len(queue)
		for i := 0; i < size; i++ {
			// Cur
			cur := queue[0]
			queue = queue[1:]
			// update: all family memebers of cur share same rank
			for _, member := range family[cur] {
				res[member/n][member%n] = rank
			}

			// Make the next move: for all family members
			for _, member := range family[cur] {
				for _, nei := range next[member] {
					degree[uf.Find(nei)]--
					if degree[uf.Find(nei)] == 0 {
						queue = append(queue, uf.Find(nei))
					}
				}
			}
		}
		rank++
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
