package leetcode

/**
 "Increment n-1 elements by 1" is equivalent to "decrementing 1 element by 1"
 So we need to find the minimum value in the array, and reduce all elements to that value - this is the min moves
*/
func minMoves(nums []int) int {
	// find the min value
	minNum := nums[0]
	for _, num := range nums {
		minNum = min(minNum, num)
	}

	// count moves
	res := 0
	for _, num := range nums {
		res += num - minNum
	}

	return res
}