package leetcode

import "strconv"

func validPrice_group(num int, k int64, x int) bool {
	// count # of bits in num
	bits := len(strconv.FormatInt(int64(num), 2))

	// counting from row 0: there are [0 : num] rows
	rows := num + 1

	price := 0
	// ith column (1-indexed): [1 : bits] columns
	for i := 1; i <= bits; i++ {
		// 2^i = group size
		// 2^(i-1) = # of 1s per group
		if i%x == 0 {
			groupOf1s := rows / powOf2(i) * powOf2(i-1)

			// rows%powOf2(i) = remainder group size
			// powOf2(i-1) = max # of 0s in remainder group
			remainder1s := max(0, rows%powOf2(i)-powOf2(i-1))

			price += groupOf1s + remainder1s
			if int64(price) > k {
				return false
			}
		}
	}
	return true
}

func powOf2(i int) int {
	return 1 << i
}
