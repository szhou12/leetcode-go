package leetcode

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// Step 1: Reconstrcut adj-list repr + Calculate degree
	n := numCourses
	degree := make([]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range prerequisites {
		from, to := edge[0], edge[1]
		next[from] = append(next[from], to)
		degree[to]++
	}

	// Step 2: Topological Sort - find all prereqs for course i
	prereqs := make([]map[int]bool, n) // prereqs[i] = all repreqs of course i
	for i := 0; i < n; i++ {
		prereqs[i] = make(map[int]bool)
		prereqs[i][i] = true // Base case: course i is prereq of itself
	}
	queue := make([]int, 0)
	// Start nodes: degree == 0
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for _, nei := range next[cur] {
			// Add all cur's prereqs to nei
			for x, _ := range prereqs[cur] {
				prereqs[nei][x] = true
			}

			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: Check if each query valid
	ans := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		if _, ok := prereqs[to][from]; ok { // if 'to' contains 'from' as prereq
			ans[i] = true
		}
	}
	return ans
}
