package amz

import "sort"

/*
Given two arrays nums1 and nums2, each of length n.
we want to make arr1 as close to arr2 as possible. i.e. nums1[i] == nums2[i] for all i.
We can apply the following operation on nums1 any number of times:
 - pick i, j where i != j, add 1 to arr1[i] and subtract 1 from nums1[j]
if we can't make arr1 exactly same as nums2, make as many elements equal as possible.

Example:
nums1 = [2, 4, 1], nums2 = [1, 2, 3]

round 1: nums1[2]+1, nums1[0]-1 => [1, 4, 2]
round 2: nums1[2]+1, nums1[1]-1 => [1, 3, 3]
Since we no longer can make nums1[1] smaller, we end up with two equal elements.

Output: 2
*/

/*
Key Observations:
寻找不变量！What is the invariant?
The operation leaves sum(nums1) unchanged!!!
This means:
1. if sum(nums1) == sum(nums2), then we can always make all n elements equal by applying the operation a certain # of times.
2. if sum(nums1) > sum(nums2), then we will have at least one pair that doesn't match (pigeonhole principle). But we can always equalize n-1 elements, and leave one element to take all excess sum(nums1) - sum(nums2).
3. if sum(nums1) < sum(nums2), this is same as Leetcode 1833. We have limited "coins" (ie. sum(nums1)) and we try to puchase as many "ice creams" (elements in nums2) as possible. The solution is greedy algorithm: make nums1 match as many of the smallest elements in nums2 as possible.

题目的operation像是：给了n个bins, 每个bin有一定数量的豆子。现在，我们需要从一个bin中取出一个豆子放到另一个bin中，反复操作，使得每个bin中豆子的数量满足nums2中的要求。
*/

func equalizeElements(nums1 []int, nums2 []int) int {
	n := len(nums1)
	sum1 := 0
	for _, v := range nums1 {
		sum1 += v
	}
	sum2 := 0
	for _, v := range nums2 {
		sum2 += v
	}

	// case 1: sum(nums1) == sum(nums2)
	if sum1 == sum2 {
		return n
	}

	// case 2: sum(nums1) > sum(nums2)
	if sum1 > sum2 {
		return n - 1
	}

	// case 3: sum(nums1) < sum(nums2)
	// ice creams = nums2, coins = sum(nums1)
	sort.Ints(nums2)
	res := 0
	for i := 0; i < n; i++ {
		if nums2[i] > sum1 {
			break
		}
		sum1 -= nums2[i]
		res++
	}
	return res
}
