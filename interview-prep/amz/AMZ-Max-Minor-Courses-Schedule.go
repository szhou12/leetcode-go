package amz

import "sort"

/*
Given two arrays of length n: major[] and minor[].
major[i] = hours to complete i-th major course
minor[i] = hours to complete i-th minor course
On each day, you only have "limit" hours to study.
You MUST study one major course per day. However, if there is time left, you can study at most one minor course on that day.
Find the maximum number of minor courses you can study after n days.

Example:
major = [4, 5, 2, 4]
minor = [5, 6, 3, 4]
limit = 7

Output: 2

One optimal scheduling can be:
Day 1: study major[0] and squeeze in minor[2]. Total hours = 4 + 3 = 7.
Day 2: study major[1]. Total hours = 5.
Day 3: study major[2] and squeeze in minor[0]. Total hours = 2 + 5 = 7.
Day 4: study major[3]. Total hours = 4.
So, you completed 2 minor courses within 4 days.
*/

/*
Key Observation: 
if we want max # of minor courses to fit, we definitely should try to math the longest minor course with the shortest major course. Because the currently longest minor course can not fit with the currently unmatched shortest major course, then it def won't fit with the later longer major courses.

Solution: Greedy (Sorting)
1. sort majors in asceding order
2. sort minors in descending order
3. mechanism: one pointer p1 for majors and another pointer p2 for minors. kepp moving p2 until major[p1] + minor[p2] <= limit, we've found a match, count++ and p1++

Time complexity = O(n\log n)
*/
func maxMinorCourses(major, minor []int, limit int) int {
	n := len(major)

	// sort major in ascending order
	sort.Slice(major, func(i, j int) bool {
		return major[i] < major[j]
	})
	// minor in descending order
	sort.Slice(minor, func(i, j int) bool {
		return minor[i] > minor[j]
	})

	minorPointer := 0 // minor pointer
	count := 0

	for i := 0; i < n; i++ {
		for minorPointer < n && major[i] + minor[minorPointer] > limit {
			minorPointer++
		}
		// update count
		if minorPointer < n {
			count++
			minorPointer++
		}
	}
	return count

}

