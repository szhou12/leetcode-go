package amz

import (
	"sort"
)

/*
Given an array of length n: nums[], and a number k.
Separate all elements in nums[] into k groups s.t. the sum of each group's median is maximized.
Note: if a group has even number of elements, the median takes the celing.

Example:
Input: nums = [1, 2, 3, 4, 5], k = 2
group 1: [5], group 2: [1,2,3,4]
Output: 5 + ceil(2.5) = 8

Input: nums = [2, 2, 1, 5, 3], k = 2
group 1: [5], group 2: [1,2,2,3]
Output: 5 + ceil(2) = 7
*/

/*
Key Oberservation:
In an array, unless it contains identical elements, the larger the array, more likely the further the median is from the max element.
Thus, we greedily put top k-1 largest elements into k-1 single-element groups. Leave the rest of small elements to form a median.

Time complexity: O(nlogn)
*/
func MaxSumKGroupsMedian(nums []int, k int) int {
	n := len(nums)

	// sort in descending order
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	res := 0

	// make first k-1 groups to be single-element groups
	// median of a single-element group is the element itself
	for i := 0; i < k-1; i++ {
		res += nums[i]
	}

	// calc k-th group's median
	start := k - 1 // start index of k-th group
	end := n - 1   // end index of k-th group
	mid := start + (end-start)/2
	m := n - (k - 1) // number of elements in k-th group
	if m%2 == 1 {
		res += nums[mid]
	} else {
		if (nums[mid]+nums[mid+1])%2 == 0 {
			res += (nums[mid] + nums[mid+1]) / 2

		} else {
			res += (nums[mid]+nums[mid+1])/2 + 1

		}
	}

	return res
}
