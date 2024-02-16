package leetcode

import (
	"math"
	"sort"
)

// The impplementation is correct BUT TLE!
func smallestRange_loop(nums [][]int) []int {
	n := len(nums)
	// i-th pointer points to the j-th element
	//   in i-th list that is cur in ordered set
	pointers := make([]int, n)
	set := make([]Pair, 0)
	for i := 0; i < n; i++ {
		set = append(set, Pair{value: nums[i][0], index: i})
	}

	var res []int
	rangeSize := math.MaxInt

	for {
		// sort set O(nlogn)
		sort.Slice(set, func(i, j int) bool {
			return set[i].value < set[j].value
		})

		curRange := set[len(set)-1].value - set[0].value
		// update range & res
		if curRange < rangeSize {
			rangeSize = curRange
			res = []int{set[0].value, set[len(set)-1].value}
		}

		i := set[0].index
		pointers[i]++
		if pointers[i] == len(nums[i]) {
			break
		}
		set = set[1:]
		set = append(set, Pair{value: nums[i][pointers[i]], index: i})
	}

	return res

}

type Pair struct {
	value int
	index int
}