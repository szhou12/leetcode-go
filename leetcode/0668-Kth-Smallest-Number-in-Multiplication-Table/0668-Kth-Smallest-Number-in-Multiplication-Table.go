package leetcode

func findKthNumber(m int, n int, k int) int {
	left := 1 * 1
	right := m * n

	for left < right {
		mid := left + (right-left)/2
		if countLessOrEqual(m, n, mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func countLessOrEqual(m, n, mid int) int {
	count := 0

	i, j := m, 1

	for i > 0 && j <= n {
		if i*j <= mid {
			count += i
			j++
		} else {
			i--
		}
	}
	return count
}
