package leetcode

import "math"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	m := len(toppingCosts)

	res := math.MaxInt
	resDiff := math.MaxInt

	for _, base := range baseCosts {
		// state = amount of each type of m toppings
		// e.g 21001 -> topping[0] 2 units, topping[1] 1 unit, topping[2] 0 units, ...
		for state := 0; state < pow(3, m); state++ {
			s := state
			topping := 0 // topping costs for current state
			for i := 0; i < m; i++ {
				topping += (s % 3) * toppingCosts[i] // 获取末位的数
				s /= 3                               // 砍掉末尾的数
			}

			if abs(base+topping-target) < resDiff {
				res = base + topping
				resDiff = abs(base + topping - target)
			} else if abs(base+topping-target) == resDiff && (base+topping) < target {
				res = base + topping
			}
		}
	}

	return res
}

func pow(a int, x int) int {
	res := math.Pow(float64(a), float64(x))
	return int(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
