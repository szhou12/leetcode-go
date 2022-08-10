package leetcode

func isIsomorphic(s string, t string) bool {
	// edge case
	if len(s) != len(t) {
		return false
	}

	dict := make(map[byte]byte)
	checkDups := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if val, ok := dict[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			if _, ok := checkDups[t[i]]; ok {
				return false
			}
			dict[s[i]] = t[i]
			checkDups[t[i]] = true
		}
	}

	return true

}
