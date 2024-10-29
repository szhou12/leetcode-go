package amz

import "sort"

/*
Given two arrays of length n: lend[] and debt[]
lend[i] := the amount of money you can lend
debt[i] := the amount of money you need to pay back the next day if you lend lend[i] today

Find the max number of days you can borrow money before you go broke (i.e. the current lending amount is not enough to pay back the previous day's debt).

Example:
lend = [4, 6, 1, 8]
debt = [7, 10, 3, 9]

Day 1: choose lend[2]=1, owe debt[2]=3, no payback in day 1
Day 2: choose lend[0]=4, owe debt[0]=7, use today's lending money to pay back previous day's debt = 4-3>0. survive!
Day 3: choose lend[3]=8, owe debt[3]=9, use today's lending money to pay back previous day's debt = 8-7>0. survive!
Day 4: choose lend[1]=6, owe debt[1]=10, use today's lending money to pay back previous day's debt = 6-9<0. broke!

Output: 3
*/


// Wrong Imeplmentation CANT pass all test cases!!!
func maxDaysBeforeBroke(lend, debt []int) int {
	n := len(lend)

	lendPairs := make([][]int, n)
	for i := 0; i < n; i++ {
		lendPairs[i] = []int{lend[i], debt[i]}
	}

	// sort: smaller debt first. if equal debt, larger lend first
	sort.Slice(lendPairs, func(i, j int) bool {
		if lendPairs[i][1] == lendPairs[j][1] {
			return lendPairs[i][0] < lendPairs[j][0]
		}
		return lendPairs[i][1] < lendPairs[j][1]
	})

	days := 0
	prevPayback := 0 // first day has no payback
	for i := 0; i < n; i++ {
		if lendPairs[i][0] - prevPayback >= 0 {
			days++
			prevPayback = lendPairs[i][1]
		} else {
			break
		}
	}
	return days
}
