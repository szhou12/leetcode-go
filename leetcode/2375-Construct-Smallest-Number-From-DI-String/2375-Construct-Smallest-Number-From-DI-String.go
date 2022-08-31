package leetcode

import "strconv"

func smallestNumber(pattern string) string {
	pattern = "I" + pattern
	n := len(pattern)
	max := 0 // NOTE: baseline starts at 0 instead of -1 bc only 1-9 allowed
	var res string

	for l := 0; l < n; l++ {
		r := l + 1
		for r < n && pattern[r] == 'D' {
			r++
		}

		count := r - l
		for k := max + count; k >= max+1; k-- {
			res = res + strconv.Itoa(k)
		}

		max = max + count
		l = r - 1

	}

	return res
}
