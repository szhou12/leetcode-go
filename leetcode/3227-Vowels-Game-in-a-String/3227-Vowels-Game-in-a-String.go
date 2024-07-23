package leetcode

func doesAliceWin(s string) bool {
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			return true
		}
	}
	return false
}