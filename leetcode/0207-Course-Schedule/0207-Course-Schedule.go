package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Step 1: pre-processing
	n := numCourses
	degree := make([]int, n)
	next := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		next[i] = make(map[int]bool)
	}
	for _, edge := range prerequisites {
		a, b := edge[0], edge[1] // a <- b
		degree[a]++
		next[b][a] = true
	}

	// Step 2: Topological sort
	count := 0
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
		// update
		count++

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: calculate result
	if count == n {
		return true
	}
	return false

}

// 这个写法不够模版化，目前已Deprecated
func canFinish_deprecated(numCourses int, prerequisites [][]int) bool {
	incoming := make([]int, numCourses)
	outgoing := make([][]int, numCourses)
	next := make([]int, 0, numCourses)

	for _, coursePair := range prerequisites {
		incoming[coursePair[0]]++
		outgoing[coursePair[1]] = append(outgoing[coursePair[1]], coursePair[0])
	}

	for i, deg := range incoming {
		if deg == 0 {
			next = append(next, i)
		}
	}

	for i := 0; i != len(next); i++ {
		course := next[i]
		outCourses := outgoing[course]
		for _, outCourse := range outCourses {
			incoming[outCourse]--
			if incoming[outCourse] == 0 {
				next = append(next, outCourse)
			}
		}
	}

	if len(next) == numCourses {
		return true
	}
	return false
}
