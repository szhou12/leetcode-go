package leetcode

// Optimal Solution: Two Pointers - 谁绝对值大移谁
// 结果从后往前添加
func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	res := make([]int, n)
	i := n - 1
	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
		i--
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 利用 MergeSort 里 merge 的思路：谁小移谁
func sortedSquares_merge(nums []int) []int {
	mid := findNonNeg(nums)
	if mid == -1 { // nums里全是负数，直接平方再逆序
		reverseSq(&nums)
		return nums
	}
	left := nums[:mid]
	reverseSq(&left)
	right := nums[mid:]
	for i := 0; i < len(right); i++ {
		right[i] = right[i] * right[i]
	}

	res := make([]int, 0)

	// Merge: 谁小移谁
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	if l < len(left) {
		res = append(res, left[l:]...)
	} else if r < len(right) {
		res = append(res, right[r:]...)
	}
	return res
}

func findNonNeg(nums []int) int {
	for i, val := range nums {
		if val >= 0 {
			return i
		}
	}
	return -1
}

func reverseSq(nums *[]int) {
	l, r := 0, len(*nums)-1
	for l <= r {
		(*nums)[l], (*nums)[r] = (*nums)[r], (*nums)[l]
		l++
		r--
	}
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] = (*nums)[i] * (*nums)[i]
	}
}
