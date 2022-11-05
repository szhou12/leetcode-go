package leetcode

func backspaceCompare(s string, t string) bool {

	newS := doBackspace(s)
	newT := doBackspace(t)
	return newS == newT
}

func doBackspace(s string) string {
	stack := make([]rune, 0)
	for _, char := range s {
		if char != '#' {
			stack = append(stack, char)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return string(stack)
}
