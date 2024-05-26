package leetcode

import (
	"strconv"
	"strings"
)

func compressedString(word string) string {
	n := len(word)

	l, r := 0, 0
	var builder strings.Builder
	for l < n {
		// when out of loop, r either 1) at n, 2) points to next different char, 3) points to same char but at 10th
		for r < n && word[l] == word[r] && r-l+1 <= 9 {
			r++
		}

		count := strconv.Itoa(r - l)
		char := string(word[l])
		builder.WriteString(count+char)

		l = r
	}

	return builder.String()
}
