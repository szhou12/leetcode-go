package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")

	// edge case
	if len(pattern) != len(sList) {
		return false
	}

	dict := make(map[byte]string)
	checkDups := make(map[string]bool)

	for i := 0; i < len(pattern); i++ {
		if str, ok := dict[pattern[i]]; ok {
			if str != sList[i] {
				return false
			}
		} else {
			if _, ok := checkDups[sList[i]]; ok {
				return false
			}
			dict[pattern[i]] = sList[i]
			checkDups[sList[i]] = true

		}
	}

	return true
}
