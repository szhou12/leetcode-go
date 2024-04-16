package leetcode

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	res := 0

	// loop k machines
	for machine := 0; machine < k; machine++ {
		res = max(res, bsMachine(budget, composition[machine], stock, cost))
	}
	return res
}

// Binary Search: find the maximum # of alloys that machine i can produce without exceeding the budget
func bsMachine(budget int, composition []int, stock []int, cost []int) int {
	left, right := 0, int(1e9)
	for left < right {
		mid := right - (right-left)/2 // total # of alloys we guess machine i can produce
		if calCost(mid, composition, stock, cost) > budget {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}

func calCost(cnt int, composition []int, stock []int, cost []int) int {
	n := len(composition)
	coins := 0
	for i := 0; i < n; i++ {
		if cnt*composition[i]-stock[i] > 0 {
			coins += cost[i] * (cnt*composition[i] - stock[i])
		}
	}
	return coins
}
