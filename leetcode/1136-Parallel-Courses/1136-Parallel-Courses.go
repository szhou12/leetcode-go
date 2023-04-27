package leetcode

func minimumSemesters(n int, relations [][]int) int {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make([]int, n+1)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range relations {
		from, to := edge[0], edge[1]
		next[from] = append(next[from], to)
		degree[to]++
	}

	// Step 2: Topological Sort
	semester := make([]int, n+1) // semester[i] = # semesters needed to study course[i]
	for i := 0; i <= n; i++ {
		semester[i] = 1
	}
	queue := make([]int, 0)
	// Start nodes: degree == 0
	for i := 1; i <= n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Loop
	visited := 0 // # of courses that can enter queue
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// udpate visited
		visited++

		// Make the next move
		for _, nei := range next[cur] {
			// udpate semester[nei]
			semester[nei] = max(semester[nei], semester[cur]+1)
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// check if has cycle
	if visited != n {
		return -1
	}

	// Step 3: find the # of semesters for the last course
	res := 0
	for i := 1; i <= n; i++ {
		res = max(res, semester[i])
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
