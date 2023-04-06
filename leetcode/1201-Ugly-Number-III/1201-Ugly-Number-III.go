package leetcode

import "math"

func nthUglyNumber(n int, a int, b int, c int) int {
	left, right := 0, math.MaxInt

	for left < right {
		mid := left + (right-left)/2

		if divisibleTotals(mid, a, b, c) < n { // [1...mid] has not enough ugly numbers
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func divisibleTotals(num, a, b, c int) int {
	setA := num / a                   // how many numbers in [1...num] that are divisible by a
	setB := num / b                   // how many numbers in [1...num] that are divisible by b
	setC := num / c                   // how many numbers in [1...num] that are divisible by c
	setAB := num / lcm(a, b)          // how many numbers in [1...num] that are divisible by a & b
	setAC := num / lcm(a, c)          // how many numbers in [1...num] that are divisible by a & c
	setBC := num / lcm(b, c)          // how many numbers in [1...num] that are divisible by b & c
	setABC := num / lcm(lcm(a, b), c) // how many numbers in [1...num] that are divisible by a & b & c
	return setA + setB + setC - setAB - setAC - setBC + setABC
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
