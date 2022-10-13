package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	for _, num1 := range nums1 {
		if _, ok := dict[num1]; !ok {
			dict[num1]++
		}
	}

	res := make([]int, 0)
	for _, num2 := range nums2 {
		if count, ok := dict[num2]; ok {
			if count == 1 {
				res = append(res, num2)
				dict[num2]--
			}
		}
	}
	return res
}
