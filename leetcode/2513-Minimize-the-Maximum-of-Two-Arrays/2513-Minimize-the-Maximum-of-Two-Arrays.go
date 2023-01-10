package leetcode

import (
	"math"

	helper "github.com/szhou12/leetcode-go/tricks"
)

// binary search
func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	left := 0
	right := math.MaxInt

	for left < right {
		mid := left + (right-left)/2
		if notLargeEnough(divisor1, divisor2, uniqueCnt1, uniqueCnt2, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func notLargeEnough(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int, n int) bool {
	a := n - n/divisor1
	b := n - n/divisor2
	c := n - (n/divisor1 + n/divisor2 - n/helper.LCM(divisor1, divisor2))

	if a < uniqueCnt1 {
		return true
	}
	if b < uniqueCnt2 {
		return true
	}
	if a+b-c < uniqueCnt1+uniqueCnt2 {
		return true
	}

	return false
}
