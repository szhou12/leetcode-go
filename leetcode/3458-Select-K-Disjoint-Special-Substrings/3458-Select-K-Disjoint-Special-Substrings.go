package leetcode

import (
	"sort"
)

func maxSubstringLength(s string, k int) bool {
	n := len(s)

	// for each letter, record its appearance indices
	pos := make([][]int, 26)
	for i, char := range s {
		c := int(char - 'a')
		pos[c] = append(pos[c], i)
	}

	// for each a-z, try to find its self-contained substring
	intervals := make([][]int, 0)

	for letter := 0; letter < 26; letter++ {
		if len(pos[letter]) == 0 {
			continue
		}

		start := pos[letter][0]
		left := start
		right := pos[letter][len(pos[letter])-1]

		valid := true

		for left <= right {
			curLetter := int(s[left] - 'a')
			right = max(right, pos[curLetter][len(pos[curLetter])-1])

			if pos[curLetter][0] < start {
				valid = false
				break
			}

			left++
		}

		if !valid {
			continue
		}
		if right-start+1 == n {
			continue
		}

		intervals = append(intervals, []int{start, right})
	}

	return count(intervals) >= k
}

func count(intervals [][]int) int {
	// sort by ending points
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	rightEnd := -1 // the first interval will def be included

	for _, interval := range intervals {
		// count if current interval's leftend > last updated rightEnd
		if interval[0] > rightEnd {
			count++
			rightEnd = interval[1]
		}
	}

	return count
}
