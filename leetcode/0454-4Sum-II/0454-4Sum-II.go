package leetcode

// My Solution
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	arr1 := allPossible(nums1, nums2)
	arr2 := allPossible(nums3, nums4)

	res := twoSum(arr1, arr2, 0)
	return res
}

func twoSum(arr1 []int, arr2 []int, sum int) int { // O(n^2) bc at most n^2 pairs from either arr1 or arr2
	dict := make(map[int]int)
	for _, e := range arr2 {
		if _, ok := dict[e]; ok {
			dict[e]++
		} else {
			dict[e] = 1
		}
	}

	res := 0
	for _, e := range arr1 {
		if val, ok := dict[sum-e]; ok {
			res += val
		}
	}
	return res
}

func allPossible(a []int, b []int) []int { // O(n^2)
	var res []int
	for _, i := range a {
		for _, j := range b {
			res = append(res, i+j)
		}
	}
	return res
}
