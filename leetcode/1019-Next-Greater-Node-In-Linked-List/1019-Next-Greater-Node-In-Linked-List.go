package leetcode

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	// Step 0: convert linked list to array
	nums := make([]int, 0)
	node := head
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}

	// Step 1: Find nextGreater
	n := len(nums)

	// monotonic decreasing stack
	stack := make([]int, 0)

	nextGreater := make([]int, n) // stores indices of next greater element

	for i := 0; i < n; i++ {
		nextGreater[i] = -1
	}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack) - 1]] < nums[i] {
			nextGreater[stack[len(stack) - 1]] = i
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}

	// Step 2: convert indices stored in nextGreater to values
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if nextGreater[i] != -1 {
			res[i] = nums[nextGreater[i]]
		}
	}

	return res

}