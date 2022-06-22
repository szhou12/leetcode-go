package leetcode

func isValid(s string) bool {
	// edge case
	// if empty string, return true
	if len(s) == 0 {
		return true
	}

	// Stack: only store left bracket
	stack := make([]rune, 0)

	for _, bracket := range s {
		if bracket == '(' || bracket == '[' || bracket == '{' {
			// 入栈
			stack = append(stack, bracket)
		} else if (bracket == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(bracket == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(bracket == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			// 弹栈
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0

}
