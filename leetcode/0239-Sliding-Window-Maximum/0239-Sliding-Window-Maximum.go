package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	dq := make([]int, 0) // store nums[] index [i-k+1 : i] (inclusive)
	res := make([]int, 0)

	for i := 0; i < n; i++ {
		// retire tail's drawfs
		for len(dq) > 0 && nums[dq[len(dq) - 1]] < nums[i] {
			dq = dq[:len(dq) - 1]
		}
		// enqueue new element
		dq = append(dq, i)
		// retire head's oldies
		for len(dq) > 0 && dq[0] <= i - k {
			dq = dq[1:]
		}

		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}

	return res
}