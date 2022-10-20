package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	i, j := 0, 0
	res := 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			res++
			i++
			j++
		} else {
			j++
		}
	}

	return res

}
