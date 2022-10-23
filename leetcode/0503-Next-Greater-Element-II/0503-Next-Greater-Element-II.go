package leetcode

// My Solution: nums + nums
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums...)

	res := make([]int, n)
	found := false
	for i := 0; i < n; i++ {
		for j := i + 1; j < i+n; j++ {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				found = true
				break
			}
		}
		if !found {
			res[i] = -1
		}
		found = false
	}
	return res
}

// Better Solution - Stack
// 套用此类型题的单调栈模版 - 从后往前看，排挤矮个子
// 用 取模 的方法 模拟 复制一倍数组的方法
func nextGreaterElements_stack(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	stack := make([]int, 0)

	// 倒着往栈里放, 就可以正着出栈, 因为我们想找第一个greater elment
	for i := 2*n - 1; i >= 0; i-- {
		// 判定个子高矮, 矮个子起开
		for len(stack) > 0 && nums[i%n] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// nums[i] 身后的第一个更大元素
		if len(stack) > 0 {
			res[i%n] = stack[len(stack)-1]
		} else {
			res[i%n] = -1
		}
		// 当前元素入栈
		stack = append(stack, nums[i%n])
	}

	return res
}
