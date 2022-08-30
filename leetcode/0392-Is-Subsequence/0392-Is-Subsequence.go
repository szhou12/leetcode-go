package leetcode

// Optimal Solution - Binary Search
func isSubsequence_BS(s string, t string) bool {
	// Step 1: pre-process t
	dict := make([][]int, 26)
	for index, char := range t {
		dict[int(char-'a')] = append(dict[int(char-'a')], index)
	}

	// Step 2: binary search
	tPointer := 0 //pointer for t
	for _, char := range s {
		// Case 1: if a char in s doesn't even exist in t
		if len(dict[int(char-'a')]) == 0 {
			return false
		}

		// Case 2: binary search to find the smallest index of this char > tPointer
		// return the index of the index
		pos := leftBound(dict[int(char-'a')], tPointer)

		if pos == -1 {
			return false
		}

		tPointer = dict[int(char-'a')][pos] + 1
	}

	return true

}

func leftBound(array []int, target int) int {
	left := 0
	right := len(array) - 1

	// 3 edge cases to check:
	// [0, 1]
	// len(array) == 1
	// left 是否越界 (超过len(array))
	for left < right { // left == right
		mid := left + (right-left)/2
		if array[mid] == target {
			right = mid
		} else if array[mid] < target {
			// if array[mid] < target: mid def NOT a solution
			left = mid + 1
		} else {
			// if array[mid] > target: mid may be a solution
			right = mid
		}
	}

	if array[left] < target {
		return -1
	}

	return left
}

// My solution
func isSubsequence(s string, t string) bool {
	// edge cases
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	l := 0

	for r := 0; r < len(t); r++ {
		if l == len(s) {
			return true
		}
		if t[r] == s[l] {
			l++
		}
	}

	return l == len(s)
}

// Same solution - better coding style
func isSubsequence_standard(s string, t string) bool {
	for len(s) > 0 && len(t) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
