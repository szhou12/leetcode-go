package leetcode

import "sort"

func maxWeight(pizzas []int) int64 {
	n := len(pizzas)
	days := n / 4

	sort.Ints(pizzas)

	p := n - 1 // point to the heaviest pizza

	res := 0

	// First, pick heaviest pizzas on odd days
	for d := 1; d <= days; d += 2 {
		res += pizzas[p]
		p--
	}

	p-- // point to currently second heaviest pizza
	//Then, pick second heaviest pizzas on even days
	for d := 2; d <= days; d += 2 {
		res += pizzas[p]
		p-=2
	}

	return int64(res)
}
