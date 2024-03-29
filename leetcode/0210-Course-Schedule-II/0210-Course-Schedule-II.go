package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	// Step 1: Reconstruct adj-list representation + Calculate degree
	n := numCourses
	degree := make([]int, n) // # of incoming edges
	next := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		next[i] = make(map[int]bool)
	}
	for _, edge := range prerequisites {
		a, b := edge[0], edge[1] // a <- b
		degree[a]++
		next[b][a] = true
	}

	// Step 2: Topological Sort
	courseOrder := make([]int, 0)
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
		courseOrder = append(courseOrder, cur)

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: calculate results
	if len(courseOrder) != n {
		return []int{}
	}
	return courseOrder

}

// 这个写法不够模版化，目前已Deprecated
func findOrder_deprecated(numCourses int, prerequisites [][]int) []int {
	// incoming[i] = # of incoming edges (pre-reqs) for course i
	incoming := make([]int, numCourses)
	// outgoing[i] = outgoing edges for course i
	outgoing := make([][]int, numCourses)
	// next = course order we want to output
	next := make([]int, 0, numCourses) // NOTE: init len=0, capacity=numCourses

	// fill up incoming & outgoing
	for _, coursePair := range prerequisites {
		incoming[coursePair[0]]++
		outgoing[coursePair[1]] = append(outgoing[coursePair[1]], coursePair[0])
	}
	// put all courses with no incoming edges (prereqs) into next first
	for i, deg := range incoming {
		if deg == 0 {
			next = append(next, i)
		}
	}

	// topological sort
	for i := 0; i != len(next); i++ {
		// get a course with no incoming edges
		course := next[i]
		// get this course's all outgoing edges
		outCourses := outgoing[course]
		for _, outCourse := range outCourses {
			incoming[outCourse]--
			if incoming[outCourse] == 0 {
				next = append(next, outCourse)
			}
		}
	}

	// If possible to finish all courses
	if len(next) == numCourses {
		return next
	}

	// Impossible to finish all courses
	return []int{}
}
