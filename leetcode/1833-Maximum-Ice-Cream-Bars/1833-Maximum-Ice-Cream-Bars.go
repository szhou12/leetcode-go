package leetcode

import "sort"

func maxIceCream(costs []int, coins int) int {
	n := len(costs)
	sort.Ints(costs)

	res := 0
	for i := 0; i < n; i++ {
		if costs[i] > coins {
			break
		}
		coins -= costs[i]
		res++
	}

	return res
}
