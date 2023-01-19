package leetcode

// diff array
func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1002)
	sum := make([]int, 1002)

	for _, trip := range trips {
		val := trip[0]
		l := trip[1]
		r := trip[2] - 1 // NOTE: 题目给出的trip是左闭右开, 因为在trip[2], 乘客已经下车。所以调整r得到双闭区间

		// update diff
		diff[l] += val
		diff[r+1] -= val
	}

	for i := 0; i < len(sum); i++ {
		if i == 0 {
			sum[i] = diff[i]
		} else {
			sum[i] = sum[i-1] + diff[i]
		}

		if sum[i] > capacity {
			return false
		}
	}

	return true

}
