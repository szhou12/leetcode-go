package leetcode

import (
	"math/rand"
	"sort"
)

/*
idx 0   1      2       3    4
w: [1,  3,     4,      2,   1]
   [0][1,2,3][4,5,6,7][8,9][10]
   [0   3      7       9    10]  keep only right-end of each interval to form nums[]
*/

type Solution struct {
	nums []int
}

func Constructor(w []int) Solution {
	nums := make([]int, 0)
	sum := -1 // start from -1 so that nums[] elements are 0-based
	for _, v := range w {
		sum += v
		nums = append(nums, sum)
	}

	return Solution{nums: nums}
}

func (this *Solution) PickIndex() int {
	maxValue := this.nums[len(this.nums)-1] + 1
	// uniformly draw a number in [0, maxValue)
	r := rand.Intn(maxValue)

	// find which interval that r falls into
	// we do so by binary search, lowerBound, find the first rightend >= r, its corresponding index is our pick
	res := sort.Search(len(this.nums), func(i int) bool {
		return this.nums[i] >= r
	})
	return res

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
