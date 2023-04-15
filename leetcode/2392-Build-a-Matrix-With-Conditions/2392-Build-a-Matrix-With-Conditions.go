package leetcode

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	// Step 1 + 2: Reconstruct + Topo Sort
	row := topo(k, rowConditions)
	col := topo(k, colConditions)

	// No solution
	if len(row) == 0 || len(col) == 0 {
		return [][]int{}
	}

	// Step 3: Fill up result matrix
	pos := make([][]int, k+1) // pos[i] = [x, y] of integer i
	for i := 0; i <= k; i++ {
		pos[i] = make([]int, 2)
	}
	for i := 0; i < len(row); i++ {
		pos[row[i]][0] = i
	}
	for j := 0; j < len(col); j++ {
		pos[col[j]][1] = j
	}

	matrix := make([][]int, k)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, k)
	}
	for i := 1; i <= k; i++ {
		x, y := pos[i][0], pos[i][1]
		matrix[x][y] = i
	}

	return matrix
}

func topo(n int, conditions [][]int) []int {
	// Step 1: Reconstruct adj-list repr + calc degree
	degree := make([]int, n+1)
	next := make([][]int, n+1)
	for i := 0; i < len(next); i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range conditions {
		a, b := edge[0], edge[1] // a -> b
		degree[b]++
		next[a] = append(next[a], b)
	}

	// Step 2: Topo Sort
	res := make([]int, 0)
	queue := make([]int, 0)
	// Start nodes: degree == 0
	for i := 1; i <= n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update
		res = append(res, cur)

		// Make the next move
		for _, nei := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Check if has cycle: when we can't order all elements
	if len(res) != n {
		return []int{}
	}

	return res
}
