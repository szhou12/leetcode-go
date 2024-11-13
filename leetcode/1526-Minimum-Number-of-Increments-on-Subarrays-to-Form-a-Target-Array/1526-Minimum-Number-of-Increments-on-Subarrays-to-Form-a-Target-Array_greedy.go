package leetcode

func minNumberOperations_greedy(target []int) int {
	n := len(target)
	res := 0
	curHeight := 0
	for i := 0; i < n; i++ {
		if target[i] > curHeight {
			res += target[i] - curHeight
		}
		curHeight = target[i]
	}

	return res
}