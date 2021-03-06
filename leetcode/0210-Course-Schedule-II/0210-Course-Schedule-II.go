package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
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
