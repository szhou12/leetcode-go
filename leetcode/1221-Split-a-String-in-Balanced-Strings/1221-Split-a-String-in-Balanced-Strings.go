package leetcode

func balancedStringSplit(s string) int {
	lCount := 0
	splits := 0

	for r := 0; r < len(s); r++ {
		if s[r] == 'L' {
			lCount++
		} else {
			lCount--
		}
		if lCount == 0 {
			splits++
		}
	}
	return splits
}
