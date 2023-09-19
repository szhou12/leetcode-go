package leetcode

func longestEqualSubarray(nums []int, k int) int {
	// Step 1: Create a map to store each number and its occuring indices
	position := make(map[int][]int)
	for idx, num := range nums {
		position[num] = append(position[num], idx)
	}

	// Step 2: 2-pointers on index array
	res := 0
	for _, pos := range position {
		right := 0
		for left := 0; left < len(pos); left++ {
			// pos[right]-pos[left]+1: total # elements in the window [left:right]
			// right-left+1: # current number in the window [left:right]
			for right < len(pos) && (pos[right]-pos[left]+1)-(right-left+1) <= k {
				res = max(res, right-left+1)
				right++
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
