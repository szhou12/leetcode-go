package leetcode

func removeDuplicates(s string) string {
	// edge case
	if len(s) == 1 {
		return s
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if len(stack) == 0 || stack[len(stack)-1] != char {
			stack = append(stack, char)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
