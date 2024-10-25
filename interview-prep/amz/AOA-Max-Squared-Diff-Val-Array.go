package amz

import "sort"

/*
 * Max Squared Difference of Values in Array
 * Heights array of length n: h where h[i] = points of i.
 * Points collected by selecting from element i to element j is (h[i] -h[j])^2.
 * The user can only collect an element once but they can choose any order to select. Find the maximal value they can collect.
 * E.g. n = 3, h = [5,2,5]. Collecting order: ground → 3rd element → 2nd element → 1st element. Max value = (0-h[2])^2 + (h[2]-h[1])^2 + (h[1] - h[0])^2 = 43
 */


// Time complexity = O(nlogn)
func maxValue(h []int) int {
	n := len(h)

	// Sort the array in increasing order
	sort.Ints(h)

	// Two pointers from opposite directions
	// 双指针相向而行
	sequence := make([]int, 0)
	left, right := 0, n-1
	for left <= right {
		if left != right {
			sequence = append(sequence, h[right])
			right--
			sequence = append(sequence, h[left])
			left++
		} else {
			sequence = append(sequence, h[left])
			left++
		}
	}

	// prepend 0
	sequence = append([]int{0}, sequence...)

	res := 0
	for i := 1; i < len(sequence); i++ {
		diff := sequence[i] - sequence[i-1]
		res += diff * diff
	}

	return res
}
