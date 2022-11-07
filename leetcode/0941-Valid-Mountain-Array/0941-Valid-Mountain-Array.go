package leetcode

func validMountainArray(arr []int) bool {
	// edge case
	if len(arr) < 3 {
		return false
	}

	l, r := 0, len(arr)-1

	for l < len(arr)-1 && arr[l] < arr[l+1] {
		l++
	}
	for r > 0 && arr[r-1] > arr[r] {
		r--
	}

	// l != r: 说明有平流层 (i.e. 左右两边至少有一边没有单调递增)
	// l == len(arr)-1: l走到头了, 说明数组一整个单调递增 e.g. [0, 1, 2, 3]
	// r == 0: r走到头了, 说明数组一整个单调递减 e.g. [3, 2, 1, 0]
	if l != r || l == len(arr)-1 || r == 0 {
		return false
	}
	return true
}
