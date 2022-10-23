package leetcode

// My Solution - naive
// 用一个map, key = nums2的元素, value = 该元素在nums2的index
// 然后再让 nums1 中的元素去查表即可
func nextGreaterElement_naive(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	for index, num := range nums2 {
		dict[num] = index
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for idx, val := range nums1 {
		startIndex := dict[val] + 1
		for i := startIndex; i < len(nums2); i++ {
			if nums2[i] > val {
				res[idx] = nums2[i]
				break // break bc we only want the first greater element
			}
		}
	}
	return res
}

// Better Solution - Stack
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greaterMap := findNextGreaterElement(nums2)

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = greaterMap[nums1[i]]
	}
	return res
}

// 此类型题的单调栈模版，只需要根据具体的题，更改返回的东西即可
func findNextGreaterElement(nums []int) map[int]int {
	res := make(map[int]int)

	stack := make([]int, 0)

	// 倒着往栈里放, 就可以正着出栈, 因为我们想找第一个greater elment
	for i := len(nums) - 1; i >= 0; i-- {
		// 判定个子高矮
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			// 矮个起开，反正也被挡着了...
			stack = stack[:len(stack)-1]
		}
		// nums[i] 身后的第一个更大元素
		if len(stack) > 0 {
			res[nums[i]] = stack[len(stack)-1]
		} else {
			res[nums[i]] = -1
		}
		// 当前元素入栈
		stack = append(stack, nums[i])
	}

	return res

}
