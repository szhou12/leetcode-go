package leetcode

func isAnagram(s string, t string) bool {
	record := make([]int, 26)
	for _, char := range s {
		idx := char - 'a'
		record[idx]++
	}

	for _, char := range t {
		idx := char - 'a'
		if record[idx] == 0 {
			return false
		}
		record[idx]--
	}

	// in case of: s = "ab", t = "a"
	for _, val := range record {
		if val != 0 {
			return false
		}
	}

	return true
}
