package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
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
