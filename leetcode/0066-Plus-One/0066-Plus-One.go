package leetcode

// Iterative
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	// Bc input digits have no leading 0, if the first digit is 0, it must mean we get a 10
	if digits[0] == 0 {
		result := append([]int{1}, digits...)
		return result
	} else {
		return digits
	}
}

// Recursion
func plusOne_recursion(digits []int) []int {
	idx := len(digits) - 1

	res := addOne(digits, idx)
	return res
}

func addOne(digits []int, index int) []int {

	if index < 0 {
		return digits
	}

	if digits[index] < 9 {
		digits[index]++
		return digits
	} else if index == 0 {
		return []int{1, 0}
	} else {
		prev := addOne(digits[:index], index-1)
		prev = append(prev, 0)
		return prev
	}
}
