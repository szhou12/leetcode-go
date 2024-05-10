package leetcode

import "math"

func findKthSmallest(coins []int, k int) int64 {
	// binary search: guess, within a given [1...mid] range, the number of values divisible by any combination of coins >= k or < k?
	// search range: coins combined values x*coins[0] + y*coins[1] + ... + z*coins[n-1]
	left, right := 1, math.MaxInt
	for left < right {
		mid := left + (right-left)/2
		if countNumbers(coins, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}

func countNumbers(coins []int, target int) int {
	n := len(coins)

	res := 0
	sign := 1
	
	for k := 1; k <= n; k++ {
		sum := 0 // total counts of values in [1...target] that are divisible by LCM of any k coins

		// Grosper's Hack: iterate over all combinations of C(n, k)
		state := (1 << k) - 1

		for state < (1 << n) {
			// LCM of cur combination of C(n, k)
			LCM := 1
			for i := 0; i < n; i++ {
				if (state >> i) & 1 == 1 {
					LCM = lcm(LCM, coins[i])
				}
			}
			sum += target / LCM // accumulate counts of values divisible by LCM of cur combination

			c := state & -state
			r := state + c
			state = (r^state)>>2/c|r
		}

		res += sum * sign // Inclusion-Exclusion Principle: + for odd k, - for even k
		sign *= -1
	}

	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
