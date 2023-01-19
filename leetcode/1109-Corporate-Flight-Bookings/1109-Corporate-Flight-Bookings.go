package leetcode

// 差分数组: 这道题的index会比较绕
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	sum := make([]int, n+1)

	// O(n)
	for _, booking := range bookings {
		// l, r 都把原来index-1, 这是因为题目里的index从1开始算, 而diff[]从0开始算
		l := booking[0] - 1
		r := booking[1] - 1
		seats := booking[2]

		diff[l] += seats
		diff[r+1] += seats // r+1不会越界因为r是原来index往前移了一位
	}

	// O(n)
	for i := 0; i < len(diff); i++ {
		if i == 0 {
			sum[0] = diff[0]
		} else {
			sum[i] = sum[i-1] + diff[i]
		}
	}

	// 截掉最后一个元素, 是多余的
	return sum[:len(sum)-1]
}
