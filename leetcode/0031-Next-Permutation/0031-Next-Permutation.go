package leetcode

func nextPermutation(nums []int) {
	// Step 1: Backwards, find the first i s.t. A[i] < A[i+1]
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Step 1b: if whole array in descending order, return its ascending order
	if i == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	// Step 2: Backwards, find the first j s.t. A[j] > A[i]
	j := len(nums) - 1
	for j >= 0 && nums[j] <= nums[i] {
		j--
	}

	// Step 3: swap A[i], A[j]
	nums[i], nums[j] = nums[j], nums[i]

	// Step 4: sort nums[i+1:] in ascending order
	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
