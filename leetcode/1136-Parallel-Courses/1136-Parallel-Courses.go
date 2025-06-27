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

// use depth and record levelSize instead of using semesters[]
func minimumSemesters_sol2(n int, relations [][]int) int {
    // step 1: convert relations to adj-list representation
    // count in-degrees of each node i
    next := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        next[i] = make([]int, 0)
    }
    inDegree := make([]int, n+1)
    for _, edge := range relations {
        from, to := edge[0], edge[1]
        next[from] = append(next[from], to)
        inDegree[to]++
    }

    depth := 0 // traversal depth is min # semesters needed
    visited := 0 // count # of nodes we visit during bfs
    queue := make([]int, 0)
    // start nodes
    for i := 1; i <= n; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
            visited++
        }
    }

    // loop
    for len(queue) > 0 {
        levelSize := len(queue)
        depth++

        for i := 0; i < levelSize; i++ {
            // pop the current
            cur := queue[0]
            queue = queue[1:]

            for _, nei := range next[cur] {
                inDegree[nei]--
                if inDegree[nei] == 0 {
                    queue = append(queue, nei)
                    visited++
                }
            }
        }
    }

    // check if has cycle
    if visited != n {
        return -1
    }

    return depth
}
