package leetcode

func secondGreaterElement(nums []int) []int {
	n := len(nums)

	// deceasing stack
	stack1 := make([]int, 0) // haven't encountered greater element yet
	stack2 := make([]int, 0) // have encountered first greater element already

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	for i := 0; i < n; i++ {
		// inspect stack2: check if incoming nums[i] is greater than any stack-top elements in stack2
		// if so, nums[i] is the second greater element for them.
		for len(stack2) > 0 && nums[stack2[len(stack2)-1]] < nums[i] {
			res[stack2[len(stack2)-1]] = nums[i]
			stack2 = stack2[:len(stack2)-1]
		}

		// inspect stack1: check if incoming nums[i] is greater than any stack-top elements in stack1
		// if so, nums[i] is the first greater element for them. We need to do 2 things:
		// 1. push those stack1-top element into stack2 in reverse order
		// 2. push i into stack1
		temp := make([]int, 0)
		for len(stack1) > 0 && nums[stack1[len(stack1)-1]] < nums[i] {
			temp = append(temp, stack1[len(stack1)-1])
			stack1 = stack1[:len(stack1)-1]
		}

		for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
			temp[i], temp[j] = temp[j], temp[i]
		}
		stack2 = append(stack2, temp...)
		
		stack1 = append(stack1, i)
	}

	return res
}
