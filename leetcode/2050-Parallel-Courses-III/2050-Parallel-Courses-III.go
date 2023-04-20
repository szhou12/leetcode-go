package leetcode

func minimumTime(n int, relations [][]int, time []int) int {
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

	// Step 2: Topological Sort, track node's 'depth' (i.e. least time to complete this course)
	t := make([]int, n+1) // t[i] = least time to complete course i
	queue := make([]int, 0)

	// Start nodes: degree == 0
	for i := 1; i <= n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
			t[i] = time[i-1]
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for _, nei := range next[cur] {
			t[nei] = max(t[nei], t[cur]+time[nei-1]) // update nei's least completion time
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: Find the course that has latest completion time
	res := 0
	for i := 0; i <= n; i++ {
		res = max(res, t[i])
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
